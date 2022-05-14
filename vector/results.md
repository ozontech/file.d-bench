# Vector.dev

## Bench 1 - Version 0.18.X alpine

### single file 500 MB

```
STATS = overall bytes: 957732258, requests: 392, timestamps:
100 MB mark took 4005 milliseconds
200 MB mark took 4133 milliseconds
300 MB mark took 4150 milliseconds
400 MB mark took 4118 milliseconds
500 MB mark took 4200 milliseconds
600 MB mark took 4068 milliseconds
700 MB mark took 4112 milliseconds
800 MB mark took 4149 milliseconds
900 MB mark took 4128 milliseconds
0 to 900 MB took 37067 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
vectordev-benchmark		CPU 201.94%	Mem 18.94MiB / 256MiB = 7.40%	NetIO 49.4kB / 49.1MB	BlockIO 27.1MB / 12.3kB
mockelasticsearch		CPU 1.28%	Mem 6.586MiB / 512MiB = 1.29%	NetIO 76MB / 71.3kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 207.36%	Mem 19.82MiB / 256MiB = 7.74%	NetIO 154kB / 169MB	BlockIO 27.1MB / 28.7kB
mockelasticsearch		CPU 1.31%	Mem 6.602MiB / 512MiB = 1.29%	NetIO 145MB / 132kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 207.89%	Mem 18.48MiB / 256MiB = 7.22%	NetIO 239kB / 265MB	BlockIO 27.1MB / 45.1kB
mockelasticsearch		CPU 1.12%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 238MB / 215kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 204.85%	Mem 17.79MiB / 256MiB = 6.95%	NetIO 319kB / 358MB	BlockIO 27.1MB / 61.4kB
mockelasticsearch		CPU 1.19%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 334MB / 297kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 200.07%	Mem 38.41MiB / 256MiB = 15.00%	NetIO 402kB / 451MB	BlockIO 27.2MB / 73.7kB
mockelasticsearch		CPU 1.39%	Mem 6.578MiB / 512MiB = 1.28%	NetIO 427MB / 379kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 199.77%	Mem 37.06MiB / 256MiB = 14.48%	NetIO 482kB / 544MB	BlockIO 27.2MB / 90.1kB
mockelasticsearch		CPU 1.56%	Mem 6.578MiB / 512MiB = 1.28%	NetIO 520MB / 459kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 194.56%	Mem 35.89MiB / 256MiB = 14.02%	NetIO 542kB / 613MB	BlockIO 27.2MB / 102kB
mockelasticsearch		CPU 1.36%	Mem 6.625MiB / 512MiB = 1.29%	NetIO 638MB / 564kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 207.97%	Mem 38.12MiB / 256MiB = 14.89%	NetIO 624kB / 706MB	BlockIO 27.2MB / 115kB
mockelasticsearch		CPU 1.21%	Mem 6.625MiB / 512MiB = 1.29%	NetIO 733MB / 646kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 197.89%	Mem 38.05MiB / 256MiB = 14.86%	NetIO 705kB / 799MB	BlockIO 27.2MB / 131kB
mockelasticsearch		CPU 1.44%	Mem 6.625MiB / 512MiB = 1.29%	NetIO 826MB / 730kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 209.66%	Mem 38.05MiB / 256MiB = 14.86%	NetIO 812kB / 920MB	BlockIO 27.2MB / 152kB
mockelasticsearch		CPU 1.14%	Mem 6.633MiB / 512MiB = 1.30%	NetIO 895MB / 790kB	BlockIO 352kB / 0B
vectordev-benchmark		CPU 0.44%	Mem 36.99MiB / 256MiB = 14.45%	NetIO 850kB / 959MB	BlockIO 27.2MB / 160kB
mockelasticsearch		CPU 0.10%	Mem 6.633MiB / 512MiB = 1.30%	NetIO 959MB / 849kB	BlockIO 352kB / 0B
```
</details>

### 500 files 1 MB each

```
STATS = overall bytes: 955660421, requests: 395, timestamps:
100 MB mark took 6370 milliseconds
200 MB mark took 4975 milliseconds
300 MB mark took 5741 milliseconds
400 MB mark took 5818 milliseconds
500 MB mark took 5718 milliseconds
600 MB mark took 5722 milliseconds
700 MB mark took 5924 milliseconds
800 MB mark took 5629 milliseconds
900 MB mark took 3638 milliseconds
0 to 900 MB took 49537 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
vectordev-benchmark		CPU 47.19%	Mem 36.45MiB / 256MiB = 14.24%	NetIO 35.3kB / 2.42MB	BlockIO 2.15MB / 4.1kB
mockelasticsearch		CPU 0.00%	Mem 1.547MiB / 512MiB = 0.30%	NetIO 1.12kB / 0B	BlockIO 0B / 0B
vectordev-benchmark		CPU 187.14%	Mem 35.75MiB / 256MiB = 13.97%	NetIO 765kB / 57.7MB	BlockIO 2.15MB / 217kB
mockelasticsearch		CPU 0.88%	Mem 6.559MiB / 512MiB = 1.28%	NetIO 30.7MB / 415kB	BlockIO 0B / 0B
vectordev-benchmark		CPU 13.77%	Mem 33.77MiB / 256MiB = 13.19%	NetIO 1.4MB / 105MB	BlockIO 2.15MB / 430kB
mockelasticsearch		CPU 0.31%	Mem 6.566MiB / 512MiB = 1.28%	NetIO 105MB / 1.4MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 172.10%	Mem 35.22MiB / 256MiB = 13.76%	NetIO 2.35MB / 179MB	BlockIO 2.15MB / 537kB
mockelasticsearch		CPU 1.22%	Mem 6.566MiB / 512MiB = 1.28%	NetIO 206MB / 2.7MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 163.86%	Mem 40.07MiB / 256MiB = 15.65%	NetIO 3.27MB / 248MB	BlockIO 2.15MB / 750kB
mockelasticsearch		CPU 0.17%	Mem 6.574MiB / 512MiB = 1.28%	NetIO 226MB / 2.98MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 118.83%	Mem 34.93MiB / 256MiB = 13.64%	NetIO 4.54MB / 346MB	BlockIO 2.15MB / 963kB
mockelasticsearch		CPU 1.25%	Mem 6.605MiB / 512MiB = 1.29%	NetIO 329MB / 4.31MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 194.20%	Mem 37.41MiB / 256MiB = 14.61%	NetIO 5.23MB / 398MB	BlockIO 2.15MB / 1.07MB
mockelasticsearch		CPU 0.94%	Mem 6.605MiB / 512MiB = 1.29%	NetIO 368MB / 4.85MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 7.64%	Mem 34.89MiB / 256MiB = 13.63%	NetIO 6.16MB / 469MB	BlockIO 2.15MB / 1.28MB
mockelasticsearch		CPU 0.88%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 466MB / 6.12MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 181.10%	Mem 37.87MiB / 256MiB = 14.79%	NetIO 6.82MB / 520MB	BlockIO 2.15MB / 1.44MB
mockelasticsearch		CPU 1.33%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 550MB / 7.2MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 5.83%	Mem 35.16MiB / 256MiB = 13.73%	NetIO 7.74MB / 589MB	BlockIO 2.15MB / 1.6MB
mockelasticsearch		CPU 0.36%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 597MB / 7.84MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 188.99%	Mem 37.43MiB / 256MiB = 14.62%	NetIO 9.17MB / 700MB	BlockIO 2.15MB / 1.81MB
mockelasticsearch		CPU 1.40%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 670MB / 8.79MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 190.08%	Mem 36.69MiB / 256MiB = 14.33%	NetIO 9.83MB / 749MB	BlockIO 2.15MB / 1.97MB
mockelasticsearch		CPU 0.45%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 720MB / 9.45MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 65.90%	Mem 36.22MiB / 256MiB = 14.15%	NetIO 10.9MB / 830MB	BlockIO 2.15MB / 2.13MB
mockelasticsearch		CPU 1.31%	Mem 6.609MiB / 512MiB = 1.29%	NetIO 820MB / 10.8MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 183.78%	Mem 39.34MiB / 256MiB = 15.37%	NetIO 11.8MB / 901MB	BlockIO 2.15MB / 2.35MB
mockelasticsearch		CPU 1.24%	Mem 6.613MiB / 512MiB = 1.29%	NetIO 871MB / 11.4MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 8.67%	Mem 34.7MiB / 256MiB = 13.55%	NetIO 12.5MB / 957MB	BlockIO 2.15MB / 2.51MB
mockelasticsearch		CPU 0.46%	Mem 6.613MiB / 512MiB = 1.29%	NetIO 955MB / 12.5MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 0.21%	Mem 34.7MiB / 256MiB = 13.55%	NetIO 12.5MB / 957MB	BlockIO 2.15MB / 2.51MB
mockelasticsearch		CPU 0.04%	Mem 6.613MiB / 512MiB = 1.29%	NetIO 957MB / 12.5MB	BlockIO 0B / 0B
vectordev-benchmark		CPU 1.44%	Mem 34.7MiB / 256MiB = 13.55%	NetIO 12.5MB / 957MB	BlockIO 2.15MB / 2.51MB
mockelasticsearch		CPU 0.04%	Mem 6.613MiB / 512MiB = 1.29%	NetIO 957MB / 12.5MB	BlockIO 0B / 0B
```
</details>


## Bench 2 - Version 0.18.X alpine

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

### 1 file 500 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 951742552, requests: 390, timestamps:
100 MB mark took 15887 milliseconds
200 MB mark took 16402 milliseconds
300 MB mark took 16492 milliseconds
400 MB mark took 16102 milliseconds
500 MB mark took 16902 milliseconds
600 MB mark took 16509 milliseconds
700 MB mark took 16694 milliseconds
800 MB mark took 16310 milliseconds
900 MB mark took 16396 milliseconds
0 to 900 MB took 147697 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 943753937, requests: 388, timestamps:
100 MB mark took 14427 milliseconds
200 MB mark took 15010 milliseconds
300 MB mark took 15607 milliseconds
400 MB mark took 14457 milliseconds
500 MB mark took 15102 milliseconds
600 MB mark took 15329 milliseconds
700 MB mark took 15880 milliseconds
800 MB mark took 15003 milliseconds
900 MB mark took 15950 milliseconds
0 to 900 MB took 136769 milliseconds
```
</details>

### 500 files 1 MB each

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 955665123, requests: 398, timestamps:
100 MB mark took 21825 milliseconds
200 MB mark took 23862 milliseconds
300 MB mark took 22114 milliseconds
400 MB mark took 20989 milliseconds
500 MB mark took 22414 milliseconds
600 MB mark took 21196 milliseconds
700 MB mark took 22990 milliseconds
800 MB mark took 22591 milliseconds
900 MB mark took 21601 milliseconds
0 to 900 MB took 199586 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 947674415, requests: 389, timestamps:
100 MB mark took 13017 milliseconds
200 MB mark took 14296 milliseconds
300 MB mark took 14285 milliseconds
400 MB mark took 14315 milliseconds
500 MB mark took 14086 milliseconds
600 MB mark took 13905 milliseconds
700 MB mark took 13807 milliseconds
800 MB mark took 13993 milliseconds
900 MB mark took 14388 milliseconds
0 to 900 MB took 126096 milliseconds
```
</details>
