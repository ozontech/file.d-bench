.PHONY: generate
generate:
	rm -f ./bench-data/*.nljson || true
	go run ./bench/generate -file-count 256 -line-count 4096

.PHONY: bench
bench:
	go run ./bench/server

.PHONY: bench-file-d
bench-file-d:
	rm ./file.d/offsets  || true
	./file.d/file.d --config ./file.d/config.yaml

.PHONY: bench-fluent-bit
bench-fluent-bit:
	./fluentbit/fluent-bit fluent-bit -c fluentbit/config.toml

.PHONY: bench-filebeat
bench-filebeat:
	rm -rf ./filebeat/data || true
	./filebeat/filebeat -c ./filebeat/config.yaml

.PHONY: bench-vector
bench-vector:
	rm ./logs/checkpoints.json || true
	./vector/vector --config ./vector/config.toml
