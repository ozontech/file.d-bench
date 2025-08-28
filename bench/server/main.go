package main

import (
	"bytes"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	addr             = flag.String("addr", "127.0.0.1:9200", "TCP address to listen to")
	debug            = flag.Bool("debug", false, "Verbose logging")
	filebeatBulkSize = flag.Int("filebeat", 50, "Filebeat bulk_max_size")
	duration         = flag.Duration("duration", time.Second*10, "Benchmark duration")
	stats            = &Stats{}
	firstReq         = true
	filebeatTemplate = []byte{123, 125}
)

var (
	MethodHead = []byte("HEAD")
	MethodGet  = []byte("GET")
	MethodPost = []byte("POST")
	MethodPut  = []byte("PUT")

	PathEmpty   = []byte("/")
	PathLicense = []byte("/_license")
	PathXpack   = []byte("/_xpack")
	PathBulk    = []byte("/_bulk")

	PathFilebeatTemplatePrefix      = []byte("/_template/filebeat-")
	PathCatFilebeatTemplatePrefix   = []byte("/_cat/templates/filebeat-")
	PathIndexTemplateFilebeatPrefix = []byte("/_index_template/filebeat-")
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
	logger.Infof("filebeat %d", *filebeatBulkSize)
	termChan := make(chan os.Signal, 2)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	srv := fasthttp.Server{
		Handler:            requestHandler,
		Name:               "default",
		MaxRequestBodySize: fasthttp.DefaultMaxRequestBodySize,
	}

	logger.Infof("started, waiting for first request")

	go func() {
		if err := srv.ListenAndServe(*addr); err != nil {
			logger.Fatalf("listen failed: %s", err)
		}
	}()

	for {
		s := <-termChan
		if s == syscall.SIGHUP {
			continue
		}

		logger.Infof("got signal to quit...")
		_ = srv.Shutdown()
		logger.Infof("bye")
		break
	}
}

type Stats struct {
	mu            sync.RWMutex
	requestsCount uint64
	overallBytes  uint64
	timestamps    []time.Time
}

func (s *Stats) WriteRequest(bytesCount uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	x := s.overallBytes / (4 * 1024 * 1024)
	if s.requestsCount == 0 {
		s.timestamps = append(s.timestamps, time.Now())
		go s.LogStats()
	}
	s.requestsCount++
	s.overallBytes += bytesCount
	y := s.overallBytes / (4 * 1024 * 1024)
	if y > x {
		s.timestamps = append(s.timestamps, time.Now())
	}
}

func (s *Stats) LogStats() {
	for {
		time.Sleep(time.Second)

		s.mu.RLock()
		x := len(s.timestamps) - 1
		if x < 0 {
			s.mu.RUnlock()
			continue
		}
		overallDuration := s.timestamps[x].Sub(s.timestamps[0]).Seconds()

		logger.Infof(
			"stats bytes=%dMb, requests=%d, throughput: %.2fMb/s",
			s.overallBytes/1024/1024,
			s.requestsCount,
			float64(s.overallBytes)/overallDuration/1024/1024,
		)
		s.mu.RUnlock()
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if *debug {
		logger.Infof("request method is %q\n", ctx.Method())
		logger.Infof("request URI is %q\n", ctx.RequestURI())
		logger.Infof("requested path is %q\n", ctx.Path())
		logger.Infof("query string is %q\n", ctx.QueryArgs())
		logger.Infof("user agent is %q\n", ctx.UserAgent())
		logger.Infof("connection has been established at %s\n", ctx.ConnTime().Format(time.RFC3339Nano))
		logger.Infof("request has been started at %s\n", ctx.Time().Format(time.RFC3339Nano))
		logger.Infof("serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	}

	method := ctx.Method()
	path := ctx.Path()

	ctx.SetContentType("application/json; charset=utf8")
	ctx.Response.Header.Set("X-elastic-product", "Elasticsearch")

	if bytes.Equal(method, MethodHead) && bytes.HasPrefix(path, PathIndexTemplateFilebeatPrefix) {
		return
	}

	// Filebeat wants to get some info
	if bytes.Equal(method, MethodGet) {
		switch {
		case bytes.Equal(path, PathEmpty):
			ctx.SetBodyString(`{
"name" : "mock",
"cluster_name" : "mock",
"cluster_uuid" : "123",
"version" : {
"number" : "8.16.0",
"build_flavor" : "default",
"build_type" : "docker",
"build_hash" : "6fc81662312141fe7691d7c1c91b8658ac17aa0d",
"build_date" : "2021-12-02T15:46:35.697268109Z",
"build_snapshot" : false,
"lucene_version" : "8.10.1",
"minimum_wire_compatibility_version" : "6.8.0",
"minimum_index_compatibility_version" : "6.0.0-beta1"
}
}`)
		case bytes.Equal(path, PathLicense):
			ctx.SetBodyString(`{
"license" : {
"status" : "active",
"uid" : "c2c7dabb-d3f6-42a3-b50b-676428c94417",
"type" : "basic",
"issue_date" : "2021-12-09T13:49:23.101Z",
"issue_date_in_millis" : 1639057763101,
"max_nodes" : 1000,
"issued_to" : "docker-cluster",
"issuer" : "elasticsearch",
"start_date_in_millis" : -1
}
}`)
		case bytes.Equal(path, PathXpack):
			ctx.SetBodyString(`
{"build":{"hash":"6fc81662312141fe7691d7c1c91b8658ac17aa0d","date":"2021-12-02T15:46:35.697268109Z"},
"license":{"uid":"c2c7dabb-d3f6-42a3-b50b-676428c94417","type":"basic","mode":"basic","status":"active"},
"features":{"aggregate_metric":{"available":true,"enabled":true},"analytics":{"available":true,"enabled":true},
"ccr":{"available":false,"enabled":true},"data_streams":{"available":true,"enabled":true},"data_tiers":{"available":true,"enabled":true},
"enrich":{"available":true,"enabled":true},"eql":{"available":true,"enabled":true},"frozen_indices":{"available":true,"enabled":true},
"graph":{"available":false,"enabled":true},"ilm":{"available":true,"enabled":true},"logstash":{"available":false,"enabled":true},
"ml":{"available":false,"enabled":true,"native_code_info":{"version":"7.16.0","build_hash":"7b1479ee4c17c2"}},
"monitoring":{"available":true,"enabled":true},"rollup":{"available":true,"enabled":true},"searchable_snapshots":{"available":false,"enabled":true},
"security":{"available":true,"enabled":false},"slm":{"available":true,"enabled":true},"spatial":{"available":true,"enabled":true},
"sql":{"available":true,"enabled":true},"transform":{"available":true,"enabled":true},"voting_only":{"available":true,"enabled":true},
"watcher":{"available":false,"enabled":true}},"tagline":"You know, for X"}`)
		case bytes.HasPrefix(path, PathCatFilebeatTemplatePrefix):
			ctx.SetBody(filebeatTemplate)
		}
		return
	}

	// Filebeat puts template
	if bytes.Equal(method, MethodPut) && bytes.HasPrefix(path, PathFilebeatTemplatePrefix) {
		filebeatTemplate = ctx.PostBody()
		ctx.SetBodyString("{\"took\":0,\"errors\":false}\n")
		return
	}

	// Finally, the bulk
	if bytes.Equal(method, MethodPost) && bytes.Equal(path, PathBulk) {
		body := ctx.PostBody()
		if ctx.IsBodyStream() {
			body := ctx.RequestBodyStream()
			buf := make([]byte, 64*1024)
			var counter uint64
			for {
				n, err := body.Read(buf)
				if err != nil {
					ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
					return
				}
				if n == 0 {
					break
				}
				counter = counter + uint64(n)
			}
			stats.WriteRequest(counter)
		} else {
			stats.WriteRequest(uint64(len(body)))
		}

		if firstReq {
			go func() {
				time.Sleep(*duration)
				dumpReport(stats)
				os.Exit(0)
			}()

			if *debug {
				msg := string(body)
				if len(body) > 1000 {
					msg = string(body[0:1000])
				}
				logger.Infof("showing %d bytes of first:\n%s", len(msg), msg)
				firstReq = false
			}
		}

		_, err := ctx.Write([]byte(`{"took":0,"errors":false,"items":[`))
		if err != nil {
			panic(err.Error())
		}
		// count should be same as (bulk_max_size - 1) in filebeat config
		for i := 1; i < *filebeatBulkSize; i++ {
			_, err = ctx.Write([]byte(`{"create": {"status": 200}},`))
			if err != nil {
				panic(err.Error())
			}
		}
		_, err = ctx.Write([]byte(`{"create":{"status": 200}}]}`))
		if err != nil {
			panic(err.Error())
		}

		return
	}

	ctx.Write([]byte("{}"))
}

func dumpReport(stats *Stats) {
	x := len(stats.timestamps) - 1
	overallDuration := stats.timestamps[x].Sub(stats.timestamps[0]).Seconds()
	logger.Infof("final result: %.2fMb/s", float64(stats.overallBytes)/overallDuration/1024/1024)
}
