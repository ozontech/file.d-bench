package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	insaneJSON "github.com/vitkovskii/insane-json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	FileName  = flag.String("file-name", "bench", "Name of the generated files")
	LineCount = flag.Int("line-count", 1000, "Lines count in each file")
	FileCount = flag.Int("file-count", 1, "Files count")
)

const (
	maxSize       = 1024 * 10
	avgSizeTarget = 1024 * 2
)

var (
	workersCount = runtime.GOMAXPROCS(0)
)

var logger = getZapLogger()

func getZapLogger() *zap.SugaredLogger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
				// TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "Instance",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		),
	).Sugar()
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano()) // constant number is intended to repeat output

	jsonFixtures, err := getJSONFixtures()
	if err != nil {
		logger.Info(err.Error())
		os.Exit(1)
	}

	fields, values := getJSONFieldsValues(jsonFixtures)

	jobs := make(chan int)

	err = os.MkdirAll("./bench-data", os.ModePerm)
	if err != nil {
		logger.Info(err.Error())
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	wg.Add(*FileCount)
	for i := 0; i < workersCount; i++ {
		go worker(*FileName, *LineCount, fields, values, jobs, wg)
	}

	go func() {
		for j := 0; j < *FileCount; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	wg.Wait()
	logger.Infof("done!")
}

func getJSONFixtures() ([]*insaneJSON.Root, error) {
	jsonFixtures := make([]*insaneJSON.Root, 0)

	//root, err = insaneJSON.DecodeFile("./fixtures/citm.json")
	//if err != nil {
	//	return nil, err
	//}
	//jsonFixtures = append(jsonFixtures, root)

	root, err := insaneJSON.DecodeFile("./bench/fixtures/twitter.json")
	if err != nil {
		return nil, err
	}
	jsonFixtures = append(jsonFixtures, root)

	root, err = insaneJSON.DecodeFile("./bench/fixtures/unknown.json")
	if err != nil {
		return nil, err
	}
	jsonFixtures = append(jsonFixtures, root)

	return jsonFixtures, nil
}

func getJSONFieldsValues(fixtures []*insaneJSON.Root) ([]string, []string) {
	fields := make([]string, 0)
	values := make([]string, 0)
	for _, root := range fixtures {
		fields, values = dfsJSONFieldsValues(root.Node, fields, values)
	}

	// deduplication
	x := make(map[string]struct{})
	for _, field := range fields {
		x[field] = struct{}{}
	}
	fields = make([]string, 0, len(x))
	for field := range x {
		fields = append(fields, field)
	}

	x = make(map[string]struct{})
	for _, value := range values {
		x[value] = struct{}{}
	}
	values = make([]string, 0, len(x))
	for value := range x {
		values = append(values, value)
	}

	return fields, values
}

func dfsJSONFieldsValues(baseNode *insaneJSON.Node, fields []string, values []string) ([]string, []string) {
	if baseNode.IsString() {
		values = append(values, string(baseNode.AsBytes()))
	}

	if baseNode.IsObject() {
		for _, node := range baseNode.AsFields() {
			fieldName := string(node.AsBytes())
			fields = append(fields, fieldName)
			fields, values = dfsJSONFieldsValues(baseNode.Dig(fieldName), fields, values)
		}
	}
	if baseNode.IsArray() {
		for _, node := range baseNode.AsArray() {
			fields, values = dfsJSONFieldsValues(node, fields, values)
		}
	}

	return fields, values
}

func worker(fileName string, lineCount int, fields []string, values []string, jobs <-chan int, wg *sync.WaitGroup) {
	root := insaneJSON.Spawn()
	defer insaneJSON.Release(root)

	buf := make([]byte, 0)
	queue := make([]*insaneJSON.Node, 0)
	jsonOut := make([]byte, 0)
	for j := range jobs {
		fullFileName := fmt.Sprintf("./bench-data/%s-%04d.nljson", fileName, j)

		curAvg := avgSizeTarget
		buf = buf[:0]
		for i := 0; i < lineCount; i++ {
			queue, jsonOut = genJSON(root, queue, jsonOut, curAvg, fields, values)
			curAvg = (curAvg + len(jsonOut)) / 2
			buf = append(buf, jsonOut...)
			buf = append(buf, '\n')
		}
		logger.Infof("file %s have been generated, avg line length is %d", fullFileName, curAvg)

		err := ioutil.WriteFile(fullFileName, buf, os.ModePerm)
		if err != nil {
			logger.Errorf(err.Error())
			os.Exit(1)
		}
		wg.Done()
	}
}

func genJSON(root *insaneJSON.Root, queue []*insaneJSON.Node, jsonOut []byte, curAvg int, fields []string, values []string) ([]*insaneJSON.Node, []byte) {
	// if we got avg size near target avg, make a jitter
	delta := float64(curAvg - avgSizeTarget)
	targetSize := 0
	if math.Abs(delta) < 200 {
		if rand.Float64() < 0.5 {
			//+ jitter
			targetSize = int((rand.Float64()*0.4 + 0.8) * float64(maxSize))
		} else {
			//- jitter
			targetSize = int((rand.Float64() * 0.2) * float64(avgSizeTarget))
		}
	} else {
		targetSize = curAvg - int(delta*rand.Float64()*0.3)
	}

	root.Clear()
	queue = queue[:0]
	l, r := 0, 0
	for {
		jsonOut = jsonOut[:0]
		jsonOut = root.Encode(jsonOut)
		if len(jsonOut) > targetSize {
			break
		}
		if l >= r {
			queue = append(queue, root.AddFieldNoAlloc(root, fields[rand.Intn(len(fields))]))
			r++
		}

		node := queue[l]
		l += 1
		nodes := genFillNodes(root, node, fields, values)
		queue = append(queue, nodes...)
		r += len(nodes)
	}

	return queue, jsonOut
}

func genFillNodes(root *insaneJSON.Root, node *insaneJSON.Node, fields []string, values []string) []*insaneJSON.Node {
	randType := func() string {
		probs := []struct {
			val float64
			t   string
		}{
			{0.2, "obj"},
			{0.2, "arr"},
			{0.3, "str"},
			{0.1, "int"},
			{0.1, "flt"},
			{0.1, "bool"},
		}
		r := rand.Float64()

		for _, x := range probs {
			if r < x.val {
				return x.t
			}
			r -= x.val
		}

		return "null"
	}

	switch randType() {
	case "obj":
		node.MutateToObject()
	case "arr":
		node.MutateToArray()
	case "str":
		node.MutateToString(values[rand.Intn(len(values))])
	case "int":
		node.MutateToInt(rand.Int())
	case "flt":
		node.MutateToFloat(rand.Float64() * 100000000)
	case "bool":
		node.MutateToBool(true)
	case "null":
		node.MutateToNull()
	}

	nodes := make([]*insaneJSON.Node, 0)
	if node.IsObject() {
		numFields := 2 + rand.Intn(5)
		for i := 0; i < numFields; i++ {
			field := fields[rand.Intn(len(fields))]
			if node.Dig(field) == nil {
				nodes = append(nodes, node.AddFieldNoAlloc(root, field))
			}
		}
	}

	if node.IsArray() {
		numFields := 2 + rand.Intn(5)
		for i := 0; i < numFields; i++ {
			nodes = append(nodes, node.AddElementNoAlloc(root))
		}
	}

	return nodes
}
