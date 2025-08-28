# Benchmarks

Pipeline: File > JSON > Elastic Search

Supported tools for benchmarking:
* File.d   0.67.0
* Vector   0.49.0
* Filebeat 9.1.2

Works only on MacOS with M CPU (Darwin ARM64) and Linux x86_64 (AMD64).

How to run:
* Run `go mod donwload`
* Unzip tools binaries `file.d/file.d_(OS).zip`, `filebeat/filebeat_(OS).zip`, `vector/vector_(OS).zip`
* Run `make generate` to generate json bench data (only once)
* Run `make bench` to run elasticsearch mock service
* (Optional) Run `make monitor-file-d|monitor-vector|monitor-filebeat` (one of commands) to monitor resource usage of specified tool via ps
* Switch to other terminal tab and run `make bench-file-d|bench-vector|bench-filebeat` (one of tool)
* Look at output of elasticsearch mock service

**Note**: if program does not shutdown with CTRL+C, use `ps | grep -E '(file.d|vector|filebeat)'`, find process starting with `./file.d|./vector|./filebeat` copy its PID and use `kill -9 PID`.

AMD Ryzen 9 2.5GHz limited to 4 cores results:

|                 |       Throughput        |   CPU usage        |   RAM usage       |
| ---             | ---                     | ---                | ---             |
| File.d   0.67.0 |     **705.19Mb/s**      | **3.75 cores**     | **301Mb**       |
| Vector   0.49.0 | &nbsp; 79.43Mb/s (-88%) |   3.59 cores (-4%) |   652Mb (+116%) |
| Filebeat 9.1.2  |       380.64Mb/s (-46%) |   3.41 cores (-9%) |   533Mb  (+77%) |
