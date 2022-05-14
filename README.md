# Benchmarks

Pipeline: File > JSON > Elastic Search 

Supported tools for benchmarking:
* File.d
* Vector
* Filebeat

How to run:
* Run `make generate` to generate json bench data (only once)
* Run `make bench` to run elasticsearch mock service
* Switch to other terminal tab and run `make bench-file-d|bench-vector|bench-filebeat` (one of tool)
* Look at output of elasticsearch mock service

2.6 GHz 6‑Core Intel Core i7 results:
* File.d — 251.48Mb/s
* Vector — 21.50Mb/s
* Filebeat — 64.75Mb/s
