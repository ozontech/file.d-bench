OS := $(shell uname | awk '{print tolower($$0)}')

.PHONY: generate
generate:
	pushd ./bench && \
	rm ./bench-data/*.nljson 2> /dev/null || true && \
	go mod download && go run ./generate.go ./logger.go -file-count 256 -line-count 4096 && \
	popd

.PHONY: bench
bench:
	go mod download && go run ./bench/bench.go ./bench/logger.go

.PHONY: bench-file-d
bench-file-d:
	pushd ./file.d && \
	rm offsets || true && \
	./file.d_$(OS) --config config.yaml && \
	popd

.PHONY: bench-fluent-bit
bench-fluent-bit:
	./fluentbit/fluent-bit fluent-bit -c fluentbit/config.toml

.PHONY: bench-filebeat
bench-filebeat:
	pushd ./filebeat && \
	rm -r data || true && \
	./filebeat_$(OS) -c config.yaml && \
	popd

.PHONY: bench-vector
bench-vector:
	pushd ./vector && \
	rm -r ./logs || true && \
	./vector_$(OS) --config config.toml && \
	popd
