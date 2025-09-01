OS := $(shell uname | awk '{print tolower($$0)}')

.PHONY: generate
generate:
	rm -f ./bench-data/*.nljson || true && \
	go run ./bench/generate -file-count 2048 -line-count 4096

.PHONY: bench
bench:
	go run ./bench/server

.PHONY: bench-file-d
bench-file-d:
	rm ./file.d/offsets || true && \
	./file.d/file.d_$(OS) --config ./file.d/config.yaml --http :9005

.PHONY: bench-filebeat
bench-filebeat:
	rm -r ./filebeat/data || true && \
	./filebeat/filebeat_$(OS) -c ./filebeat/config.yaml

.PHONY: bench-vector
bench-vector:
	rm -r ./vector/logs || true && \
	./vector/vector_$(OS) --config ./vector/config.toml

.PHONY: monitor-file-d
monitor-file-d:
	resources_monitor/monitor.sh 'file.d' '(bench|rm|monitor)'

.PHONY: monitor-filebeat
monitor-filebeat:
	resources_monitor/monitor.sh 'filebeat' '(bench|rm|monitor)'

.PHONY: monitor-vector
monitor-vector:
	resources_monitor/monitor.sh 'vector' '(bench|rm|monitor)'
