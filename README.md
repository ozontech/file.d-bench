# Benchmarks

Pipeline: File > JSON > Elastic Search

Supported tools for benchmarking:

* File.d v0.25.2
* Vector v0.38.0
* Filebeat 7.17.13

Works only on x86-64 MacOS/Linux.

How to run:

* Download a log collector binary to its folder
* Run `make generate` to generate json bench data (only once)
* Run `make bench` to run elasticsearch mock service
* Switch to other terminal tab and run `make bench-file-d|bench-vector|bench-filebeat` (one of tool)
* Look at output of elasticsearch mock service

AMD Ryzen 9 5950x (4.9 GHz 16‑Core) and SSD M2 Samsung 980 PRO (read 7000 MB/s, write 5000 MB/s) results:

* File.d — 500.48Mb/s
* Vector — 323.14Mb/s
* Filebeat — 75.42Mb/s
