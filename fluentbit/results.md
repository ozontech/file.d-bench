# Fluent-bit

## Bench 1 - Version 1.8

### single file 500 MB

```
STATS = overall bytes: 785759725, requests: 275, timestamps:
100 MB mark took 493 milliseconds
200 MB mark took 55 milliseconds
300 MB mark took 2581 milliseconds
400 MB mark took 6872 milliseconds
500 MB mark took 271 milliseconds
600 MB mark took 2926 milliseconds
700 MB mark took 4878 milliseconds
0 to 700 MB took 18079 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 0.00%	Mem 1.473MiB / 512MiB = 0.29%	NetIO 952B / 0B	BlockIO 77.8kB / 0B
fluentbit-benchmark		CPU 15.46%	Mem 106.8MiB / 256MiB = 41.74%	NetIO 1.15kB / 0B	BlockIO 0B / 0B
mockelasticsearch		CPU 13.95%	Mem 182.2MiB / 512MiB = 35.60%	NetIO 4.78MB / 29.5kB	BlockIO 143kB / 0B
fluentbit-benchmark		CPU 89.86%	Mem 255.9MiB / 256MiB = 99.97%	NetIO 50.5kB / 8.89MB	BlockIO 64.3MB / 219MB
mockelasticsearch		CPU 1.64%	Mem 365.8MiB / 512MiB = 71.44%	NetIO 264MB / 611kB	BlockIO 143kB / 0B
fluentbit-benchmark		CPU 16.45%	Mem 98.16MiB / 256MiB = 38.34%	NetIO 561kB / 242MB	BlockIO 312MB / 387MB
mockelasticsearch		CPU 0.02%	Mem 357.6MiB / 512MiB = 69.85%	NetIO 346MB / 727kB	BlockIO 143kB / 0B
fluentbit-benchmark		CPU 18.46%	Mem 167.5MiB / 256MiB = 65.44%	NetIO 729kB / 346MB	BlockIO 317MB / 387MB
mockelasticsearch		CPU 1.64%	Mem 364.8MiB / 512MiB = 71.25%	NetIO 370MB / 795kB	BlockIO 143kB / 0B
fluentbit-benchmark		CPU 51.01%	Mem 60.43MiB / 256MiB = 23.60%	NetIO 1.18MB / 584MB	BlockIO 554MB / 725MB
mockelasticsearch		CPU 5.79%	Mem 357.8MiB / 512MiB = 69.89%	NetIO 696MB / 1.31MB	BlockIO 217kB / 0B
fluentbit-benchmark		CPU 71.34%	Mem 193.9MiB / 256MiB = 75.73%	NetIO 1.2MB / 604MB	BlockIO 612MB / 725MB
mockelasticsearch		CPU 0.91%	Mem 358.9MiB / 512MiB = 70.11%	NetIO 706MB / 1.33MB	BlockIO 217kB / 0B
fluentbit-benchmark		CPU 0.02%	Mem 96.95MiB / 256MiB = 37.87%	NetIO 1.31MB / 696MB	BlockIO 612MB / 725MB
mockelasticsearch		CPU 0.02%	Mem 357.8MiB / 512MiB = 69.89%	NetIO 787MB / 1.45MB	BlockIO 217kB / 0B
fluentbit-benchmark		CPU 0.03%	Mem 97.64MiB / 256MiB = 38.14%	NetIO 1.45MB / 787MB	BlockIO 614MB / 725MB
```
</details>

### 500 files 1 MB each

256 MB of memory is not enough for fluent-bit in this bench - it hits OOM. 512MB is ok.

```
STATS = overall bytes: 778028117, requests: 272, timestamps:
100 MB mark took 3197 milliseconds
200 MB mark took 15603 milliseconds
300 MB mark took 2706 milliseconds
400 MB mark took 699 milliseconds
500 MB mark took 3734 milliseconds
600 MB mark took 1825 milliseconds
700 MB mark took 2786 milliseconds
0 to 700 MB took 30551 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 0.00%	Mem 1.605MiB / 512MiB = 0.31%	NetIO 1.01kB / 0B	BlockIO 0B / 0B
fluentbit-benchmark		CPU 6.59%	Mem 67.75MiB / 512MiB = 13.23%	NetIO 1.01kB / 0B	BlockIO 0B / 0B
mockelasticsearch		CPU 0.00%	Mem 1.605MiB / 512MiB = 0.31%	NetIO 1.08kB / 0B	BlockIO 0B / 0B
fluentbit-benchmark		CPU 77.20%	Mem 295.9MiB / 512MiB = 57.80%	NetIO 24.8kB / 3.67MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.00%	Mem 285.3MiB / 512MiB = 55.73%	NetIO 7.55MB / 43.9kB	BlockIO 49.2kB / 0B
fluentbit-benchmark		CPU 19.89%	Mem 452.9MiB / 512MiB = 88.46%	NetIO 226kB / 89.2MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.00%	Mem 310.7MiB / 512MiB = 60.68%	NetIO 120MB / 311kB	BlockIO 49.2kB / 0B
fluentbit-benchmark		CPU 11.06%	Mem 483.4MiB / 512MiB = 94.41%	NetIO 312kB / 120MB	BlockIO 0B / 16.4kB
mockelasticsearch		CPU 2.82%	Mem 473MiB / 512MiB = 92.39%	NetIO 189MB / 594kB	BlockIO 2.25MB / 0B
fluentbit-benchmark		CPU 13.54%	Mem 498.4MiB / 512MiB = 97.34%	NetIO 424kB / 156MB	BlockIO 3.99MB / 141MB
mockelasticsearch		CPU 3.43%	Mem 509.5MiB / 512MiB = 99.52%	NetIO 230MB / 826kB	BlockIO 4.69MB / 61.8MB
fluentbit-benchmark		CPU 16.38%	Mem 481MiB / 512MiB = 93.94%	NetIO 762kB / 220MB	BlockIO 34.6MB / 207MB
mockelasticsearch		CPU 0.10%	Mem 509.8MiB / 512MiB = 99.57%	NetIO 258MB / 910kB	BlockIO 10.6MB / 70.8MB
fluentbit-benchmark		CPU 14.52%	Mem 500.3MiB / 512MiB = 97.72%	NetIO 911kB / 258MB	BlockIO 65.9MB / 277MB
mockelasticsearch		CPU 11.08%	Mem 483MiB / 512MiB = 94.34%	NetIO 377MB / 1.26MB	BlockIO 21.1MB / 122MB
fluentbit-benchmark		CPU 14.21%	Mem 491.6MiB / 512MiB = 96.01%	NetIO 1.11MB / 322MB	BlockIO 80.8MB / 381MB
mockelasticsearch		CPU 0.05%	Mem 503.6MiB / 512MiB = 98.36%	NetIO 432MB / 1.42MB	BlockIO 27.1MB / 136MB
fluentbit-benchmark		CPU 17.78%	Mem 491MiB / 512MiB = 95.89%	NetIO 1.42MB / 432MB	BlockIO 132MB / 455MB
mockelasticsearch		CPU 5.59%	Mem 509.6MiB / 512MiB = 99.54%	NetIO 520MB / 1.81MB	BlockIO 44.3MB / 136MB
fluentbit-benchmark		CPU 12.19%	Mem 493.3MiB / 512MiB = 96.35%	NetIO 1.65MB / 484MB	BlockIO 161MB / 551MB
mockelasticsearch		CPU 1.21%	Mem 510.4MiB / 512MiB = 99.69%	NetIO 564MB / 2.03MB	BlockIO 48.8MB / 177MB
fluentbit-benchmark		CPU 10.81%	Mem 475.7MiB / 512MiB = 92.91%	NetIO 2.07MB / 575MB	BlockIO 224MB / 609MB
mockelasticsearch		CPU 10.95%	Mem 490.9MiB / 512MiB = 95.89%	NetIO 735MB / 3.11MB	BlockIO 68MB / 201MB
fluentbit-benchmark		CPU 0.06%	Mem 181.5MiB / 512MiB = 35.45%	NetIO 3.11MB / 735MB	BlockIO 264MB / 621MB
mockelasticsearch		CPU 0.02%	Mem 492.6MiB / 512MiB = 96.21%	NetIO 778MB / 3.18MB	BlockIO 69.7MB / 201MB
fluentbit-benchmark		CPU 0.06%	Mem 179.2MiB / 512MiB = 35.00%	NetIO 3.18MB / 778MB	BlockIO 300MB / 621MB
mockelasticsearch		CPU 0.09%	Mem 492.6MiB / 512MiB = 96.21%	NetIO 778MB / 3.18MB	BlockIO 69.7MB / 201MB
fluentbit-benchmark		CPU 0.28%	Mem 179.2MiB / 512MiB = 35.00%	NetIO 3.18MB / 778MB	BlockIO 300MB / 621MB
mockelasticsearch		CPU 0.05%	Mem 493.1MiB / 512MiB = 96.31%	NetIO 781MB / 3.18MB	BlockIO 70.2MB / 201MB
fluentbit-benchmark		CPU 3.78%	Mem 181.1MiB / 512MiB = 35.38%	NetIO 3.18MB / 781MB	BlockIO 302MB / 621MB
```
</details>

## Bench 2 - Version 1.8

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

On Linux server fluent-bit hits OOM, so memory increased up to 2048 MB and for `mockelasticsearch` too.

### 1 file 500 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 777759691, requests: 272, timestamps:
100 MB mark took 7 milliseconds
200 MB mark took 5 milliseconds
300 MB mark took 2833 milliseconds
400 MB mark took 6469 milliseconds
500 MB mark took 2 milliseconds
600 MB mark took 4016 milliseconds
700 MB mark took 4396 milliseconds
0 to 700 MB took 17731 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 744092612, requests: 256, timestamps:
100 MB mark took 105 milliseconds
200 MB mark took 70 milliseconds
300 MB mark took 10 milliseconds
400 MB mark took 9 milliseconds
500 MB mark took 61 milliseconds
600 MB mark took 12 milliseconds
700 MB mark took 203 milliseconds
0 to 700 MB took 472 milliseconds
```
</details>

### 500 files 1 MB each

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 778027803, requests: 272, timestamps:
100 MB mark took 4407 milliseconds
200 MB mark took 9074 milliseconds
300 MB mark took 1488 milliseconds
400 MB mark took 7633 milliseconds
500 MB mark took 1394 milliseconds
600 MB mark took 3087 milliseconds
700 MB mark took 1285 milliseconds
0 to 700 MB took 28370 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 741308055, requests: 255, timestamps:
100 MB mark took 28 milliseconds
200 MB mark took 71 milliseconds
300 MB mark took 19 milliseconds
400 MB mark took 20 milliseconds
500 MB mark took 64 milliseconds
600 MB mark took 11 milliseconds
700 MB mark took 198 milliseconds
0 to 700 MB took 414 milliseconds
```
</details>
