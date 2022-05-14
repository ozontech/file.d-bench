# Filebeat

## Bench 1 - Version 7.15.2

### single file 500 MB

#### bulk = 50 (default)

```
STATS = overall bytes: 1449337713, requests: 40040, timestamps:
100 MB mark took 2326 milliseconds
200 MB mark took 2353 milliseconds
300 MB mark took 2262 milliseconds
400 MB mark took 2372 milliseconds
500 MB mark took 2343 milliseconds
600 MB mark took 2290 milliseconds
700 MB mark took 2345 milliseconds
800 MB mark took 2294 milliseconds
900 MB mark took 2340 milliseconds
1000 MB mark took 2359 milliseconds
1100 MB mark took 2382 milliseconds
1200 MB mark took 2429 milliseconds
1300 MB mark took 2302 milliseconds
0 to 1300 MB took 30403 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 28.98%	Mem 2.328MiB / 512MiB = 0.45%	NetIO 82MB / 4.5MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 159.25%	Mem 116.8MiB / 256MiB = 45.61%	NetIO 1.98MB / 36MB	BlockIO 20.4MB / 8.19kB
mockelasticsearch		CPU 30.09%	Mem 2.363MiB / 512MiB = 0.46%	NetIO 204MB / 11.2MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 170.88%	Mem 71.9MiB / 256MiB = 28.09%	NetIO 13.9MB / 253MB	BlockIO 20.4MB / 1.52MB
mockelasticsearch		CPU 30.34%	Mem 2.391MiB / 512MiB = 0.47%	NetIO 423MB / 23.1MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 169.14%	Mem 71.12MiB / 256MiB = 27.78%	NetIO 20.6MB / 377MB	BlockIO 20.4MB / 1.52MB
mockelasticsearch		CPU 29.29%	Mem 2.41MiB / 512MiB = 0.47%	NetIO 547MB / 29.9MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 166.96%	Mem 72.48MiB / 256MiB = 28.31%	NetIO 32.4MB / 593MB	BlockIO 20.4MB / 1.52MB
mockelasticsearch		CPU 28.51%	Mem 2.41MiB / 512MiB = 0.47%	NetIO 764MB / 41.8MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 166.51%	Mem 71.92MiB / 256MiB = 28.09%	NetIO 39.2MB / 717MB	BlockIO 20.4MB / 1.52MB
mockelasticsearch		CPU 29.01%	Mem 2.426MiB / 512MiB = 0.47%	NetIO 936MB / 51.1MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 165.10%	Mem 74.52MiB / 256MiB = 29.11%	NetIO 48.6MB / 890MB	BlockIO 20.4MB / 1.52MB
mockelasticsearch		CPU 29.50%	Mem 2.426MiB / 512MiB = 0.47%	NetIO 1.06GB / 57.8MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 173.81%	Mem 75.55MiB / 256MiB = 29.51%	NetIO 60.3MB / 1.1GB	BlockIO 20.4MB / 1.52MB
mockelasticsearch		CPU 28.79%	Mem 2.426MiB / 512MiB = 0.47%	NetIO 1.27GB / 69.3MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 170.01%	Mem 78.88MiB / 256MiB = 30.81%	NetIO 66.9MB / 1.23GB	BlockIO 20.5MB / 1.52MB
mockelasticsearch		CPU 29.18%	Mem 2.43MiB / 512MiB = 0.47%	NetIO 1.44GB / 78.7MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 169.91%	Mem 73.7MiB / 256MiB = 28.79%	NetIO 76.1MB / 1.39GB	BlockIO 20.8MB / 1.72MB
mockelasticsearch		CPU 1.48%	Mem 2.383MiB / 512MiB = 0.47%	NetIO 1.48GB / 81MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 0.10%	Mem 72.68MiB / 256MiB = 28.39%	NetIO 81MB / 1.48GB	BlockIO 20.8MB / 2.63MB
```
</details>

#### bulk size = 1000

```
STATS = overall bytes: 1449337713, requests: 2930, timestamps:
100 MB mark took 1750 milliseconds
200 MB mark took 1666 milliseconds
300 MB mark took 1708 milliseconds
400 MB mark took 1693 milliseconds
500 MB mark took 1631 milliseconds
600 MB mark took 1730 milliseconds
700 MB mark took 1754 milliseconds
800 MB mark took 1660 milliseconds
900 MB mark took 1680 milliseconds
1000 MB mark took 1690 milliseconds
1100 MB mark took 1676 milliseconds
1200 MB mark took 1678 milliseconds
1300 MB mark took 1748 milliseconds
0 to 1300 MB took 22069 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
filebeat-benchmark		CPU 192.44%	Mem 112.4MiB / 256MiB = 43.92%	NetIO 11.6MB / 169MB	BlockIO 22.9MB / 8.19kB
mockelasticsearch		CPU 15.34%	Mem 3.797MiB / 512MiB = 0.74%	NetIO 107MB / 7.33MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 181.32%	Mem 81.52MiB / 256MiB = 31.84%	NetIO 23MB / 336MB	BlockIO 22.9MB / 8.19kB
mockelasticsearch		CPU 15.47%	Mem 3.586MiB / 512MiB = 0.70%	NetIO 401MB / 27.4MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 181.79%	Mem 71.38MiB / 256MiB = 27.88%	NetIO 43.4MB / 636MB	BlockIO 22.9MB / 336kB
mockelasticsearch		CPU 16.02%	Mem 3.586MiB / 512MiB = 0.70%	NetIO 571MB / 39MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 178.56%	Mem 71.89MiB / 256MiB = 28.08%	NetIO 54.5MB / 798MB	BlockIO 22.9MB / 336kB
mockelasticsearch		CPU 16.50%	Mem 3.781MiB / 512MiB = 0.74%	NetIO 863MB / 59MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 187.20%	Mem 71.72MiB / 256MiB = 28.02%	NetIO 70MB / 1.02GB	BlockIO 22.9MB / 336kB
mockelasticsearch		CPU 16.18%	Mem 3.691MiB / 512MiB = 0.72%	NetIO 1.09GB / 74.3MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 180.34%	Mem 72.85MiB / 256MiB = 28.46%	NetIO 89.4MB / 1.31GB	BlockIO 23.1MB / 336kB
mockelasticsearch		CPU 14.82%	Mem 3.66MiB / 512MiB = 0.71%	NetIO 1.25GB / 85.3MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 4.71%	Mem 71.89MiB / 256MiB = 28.08%	NetIO 101MB / 1.47GB	BlockIO 23.1MB / 336kB
mockelasticsearch		CPU 15.82%	Mem 4.176MiB / 512MiB = 0.82%	NetIO 1.47GB / 101MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 0.14%	Mem 71.89MiB / 256MiB = 28.08%	NetIO 101MB / 1.47GB	BlockIO 23.1MB / 336kB
mockelasticsearch		CPU 0.10%	Mem 3.66MiB / 512MiB = 0.71%	NetIO 1.47GB / 101MB	BlockIO 0B / 0B
```
</details>

### 500 files 1 MB each

#### bulk = 50 (default)

```
STATS = overall bytes: 1441779901, requests: 40060, timestamps:
100 MB mark took 4343 milliseconds
200 MB mark took 3743 milliseconds
300 MB mark took 3145 milliseconds
400 MB mark took 3065 milliseconds
500 MB mark took 3053 milliseconds
600 MB mark took 2857 milliseconds
700 MB mark took 3009 milliseconds
800 MB mark took 3225 milliseconds
900 MB mark took 3743 milliseconds
1000 MB mark took 3814 milliseconds
1100 MB mark took 3697 milliseconds
1200 MB mark took 3447 milliseconds
1300 MB mark took 2672 milliseconds
0 to 1300 MB took 43819 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
filebeat-benchmark		CPU 194.62%	Mem 122MiB / 256MiB = 47.65%	NetIO 2.31MB / 40MB	BlockIO 22.3MB / 49.2kB
mockelasticsearch		CPU 12.36%	Mem 2.371MiB / 512MiB = 0.46%	NetIO 16.1MB / 927kB	BlockIO 0B / 0B
filebeat-benchmark		CPU 201.80%	Mem 105.5MiB / 256MiB = 41.21%	NetIO 7.88MB / 137MB	BlockIO 22.4MB / 5.36MB
mockelasticsearch		CPU 17.42%	Mem 2.461MiB / 512MiB = 0.48%	NetIO 108MB / 6.23MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 220.81%	Mem 113.3MiB / 256MiB = 44.26%	NetIO 14.3MB / 249MB	BlockIO 22.4MB / 6.18MB
mockelasticsearch		CPU 18.57%	Mem 2.414MiB / 512MiB = 0.47%	NetIO 215MB / 12.3MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 222.64%	Mem 122.3MiB / 256MiB = 47.78%	NetIO 19.6MB / 343MB	BlockIO 22.4MB / 6.73MB
mockelasticsearch		CPU 20.10%	Mem 2.477MiB / 512MiB = 0.48%	NetIO 377MB / 21.5MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 223.87%	Mem 135.2MiB / 256MiB = 52.80%	NetIO 29MB / 508MB	BlockIO 22.4MB / 7.85MB
mockelasticsearch		CPU 19.99%	Mem 2.5MiB / 512MiB = 0.49%	NetIO 470MB / 26.8MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 222.07%	Mem 143.8MiB / 256MiB = 56.18%	NetIO 36.8MB / 644MB	BlockIO 22.4MB / 8.82MB
mockelasticsearch		CPU 19.41%	Mem 2.512MiB / 512MiB = 0.49%	NetIO 606MB / 34.6MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 221.08%	Mem 146.8MiB / 256MiB = 57.35%	NetIO 44.2MB / 775MB	BlockIO 22.5MB / 9.66MB
mockelasticsearch		CPU 17.90%	Mem 2.465MiB / 512MiB = 0.48%	NetIO 741MB / 42.3MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 215.77%	Mem 152.4MiB / 256MiB = 59.52%	NetIO 49.1MB / 862MB	BlockIO 22.5MB / 10.4MB
mockelasticsearch		CPU 16.37%	Mem 2.52MiB / 512MiB = 0.49%	NetIO 891MB / 50.8MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 204.71%	Mem 166.2MiB / 256MiB = 64.91%	NetIO 55.2MB / 968MB	BlockIO 23.1MB / 11.3MB
mockelasticsearch		CPU 16.35%	Mem 2.531MiB / 512MiB = 0.49%	NetIO 998MB / 56.9MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 206.51%	Mem 184MiB / 256MiB = 71.89%	NetIO 61.1MB / 1.07GB	BlockIO 23.2MB / 12MB
mockelasticsearch		CPU 17.11%	Mem 2.633MiB / 512MiB = 0.51%	NetIO 1.1GB / 62.7MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 210.13%	Mem 186.7MiB / 256MiB = 72.94%	NetIO 68.8MB / 1.21GB	BlockIO 23.2MB / 19.1MB
mockelasticsearch		CPU 17.90%	Mem 2.531MiB / 512MiB = 0.49%	NetIO 1.18GB / 67.1MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 225.50%	Mem 185.5MiB / 256MiB = 72.46%	NetIO 75.9MB / 1.33GB	BlockIO 23.4MB / 23.4MB
mockelasticsearch		CPU 18.35%	Mem 2.539MiB / 512MiB = 0.50%	NetIO 1.29GB / 73.8MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 121.63%	Mem 183.3MiB / 256MiB = 71.62%	NetIO 84.1MB / 1.48GB	BlockIO 23.4MB / 25.6MB
mockelasticsearch		CPU 22.10%	Mem 2.492MiB / 512MiB = 0.49%	NetIO 1.45GB / 82.7MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 9.46%	Mem 183.3MiB / 256MiB = 71.59%	NetIO 84.1MB / 1.48GB	BlockIO 23.4MB / 25.6MB
mockelasticsearch		CPU 0.09%	Mem 2.492MiB / 512MiB = 0.49%	NetIO 1.48GB / 84.1MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 0.68%	Mem 183.3MiB / 256MiB = 71.60%	NetIO 84.1MB / 1.48GB	BlockIO 23.4MB / 25.6MB
mockelasticsearch		CPU 0.05%	Mem 2.492MiB / 512MiB = 0.49%	NetIO 1.48GB / 84.1MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 1.91%	Mem 183.3MiB / 256MiB = 71.60%	NetIO 84.1MB / 1.48GB	BlockIO 23.4MB / 25.6MB
mockelasticsearch		CPU 0.04%	Mem 2.492MiB / 512MiB = 0.49%	NetIO 1.48GB / 84.1MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 6.02%	Mem 183.8MiB / 256MiB = 71.80%	NetIO 84.1MB / 1.48GB	BlockIO 23.4MB / 25.6MB
mockelasticsearch		CPU 0.04%	Mem 2.492MiB / 512MiB = 0.49%	NetIO 1.48GB / 84.1MB	BlockIO 0B / 0B
```
</details>

#### bulk = 1000

```
STATS = overall bytes: 1441779901, requests: 2932, timestamps:
100 MB mark took 2550 milliseconds
200 MB mark took 2706 milliseconds
300 MB mark took 2742 milliseconds
400 MB mark took 2293 milliseconds
500 MB mark took 2333 milliseconds
600 MB mark took 2301 milliseconds
700 MB mark took 2170 milliseconds
800 MB mark took 2277 milliseconds
900 MB mark took 2574 milliseconds
1000 MB mark took 2563 milliseconds
1100 MB mark took 2695 milliseconds
1200 MB mark took 2483 milliseconds
1300 MB mark took 2047 milliseconds
0 to 1300 MB took 31739 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 12.43%	Mem 3.66MiB / 512MiB = 0.71%	NetIO 55.8MB / 3.83MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 128.07%	Mem 138MiB / 256MiB = 53.89%	NetIO 925kB / 13.8MB	BlockIO 21.9MB / 8.19kB
mockelasticsearch		CPU 12.43%	Mem 3.664MiB / 512MiB = 0.72%	NetIO 168MB / 11.6MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 209.66%	Mem 109.4MiB / 256MiB = 42.74%	NetIO 14.3MB / 206MB	BlockIO 22.1MB / 4.31MB
mockelasticsearch		CPU 12.99%	Mem 3.664MiB / 512MiB = 0.72%	NetIO 309MB / 21.4MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 214.19%	Mem 119MiB / 256MiB = 46.49%	NetIO 24.5MB / 353MB	BlockIO 22.1MB / 4.85MB
mockelasticsearch		CPU 12.15%	Mem 3.727MiB / 512MiB = 0.73%	NetIO 524MB / 36.4MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 220.63%	Mem 129.4MiB / 256MiB = 50.55%	NetIO 33MB / 476MB	BlockIO 22.1MB / 5.27MB
mockelasticsearch		CPU 12.86%	Mem 3.742MiB / 512MiB = 0.73%	NetIO 701MB / 48.7MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 226.38%	Mem 142.5MiB / 256MiB = 55.66%	NetIO 45.1MB / 651MB	BlockIO 22.1MB / 5.82MB
mockelasticsearch		CPU 12.87%	Mem 4.102MiB / 512MiB = 0.80%	NetIO 873MB / 60.7MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 226.36%	Mem 152.7MiB / 256MiB = 59.65%	NetIO 57.5MB / 828MB	BlockIO 22.2MB / 6.52MB
mockelasticsearch		CPU 11.98%	Mem 3.789MiB / 512MiB = 0.74%	NetIO 1.03GB / 71.3MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 214.65%	Mem 164.8MiB / 256MiB = 64.38%	NetIO 68.5MB / 984MB	BlockIO 22.2MB / 7.08MB
mockelasticsearch		CPU 11.68%	Mem 3.781MiB / 512MiB = 0.74%	NetIO 1.18GB / 81.8MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 215.12%	Mem 177.3MiB / 256MiB = 69.24%	NetIO 78.8MB / 1.13GB	BlockIO 22.3MB / 7.63MB
mockelasticsearch		CPU 12.88%	Mem 3.938MiB / 512MiB = 0.77%	NetIO 1.34GB / 93.5MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 220.76%	Mem 188.7MiB / 256MiB = 73.69%	NetIO 89.8MB / 1.29GB	BlockIO 22.3MB / 8.19MB
mockelasticsearch		CPU 3.91%	Mem 3.793MiB / 512MiB = 0.74%	NetIO 1.47GB / 102MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 10.63%	Mem 178.3MiB / 256MiB = 69.64%	NetIO 102MB / 1.47GB	BlockIO 22.3MB / 18.7MB
mockelasticsearch		CPU 0.04%	Mem 3.793MiB / 512MiB = 0.74%	NetIO 1.47GB / 102MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 6.15%	Mem 178.3MiB / 256MiB = 69.65%	NetIO 102MB / 1.47GB	BlockIO 22.3MB / 18.7MB
mockelasticsearch		CPU 0.07%	Mem 3.793MiB / 512MiB = 0.74%	NetIO 1.47GB / 102MB	BlockIO 0B / 0B
filebeat-benchmark		CPU 1.14%	Mem 178.3MiB / 256MiB = 69.65%	NetIO 102MB / 1.47GB	BlockIO 22.3MB / 18.7MB
```
</details>


## Bench 2 - Version 7.15.2

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

### 1 file 500 MB

#### bulk = 50 (default)

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 1443331150, requests: 40040, timestamps:
100 MB mark took 7310 milliseconds
200 MB mark took 7384 milliseconds
300 MB mark took 7322 milliseconds
400 MB mark took 7299 milliseconds
500 MB mark took 7298 milliseconds
600 MB mark took 7389 milliseconds
700 MB mark took 7202 milliseconds
800 MB mark took 7310 milliseconds
900 MB mark took 7382 milliseconds
1000 MB mark took 7316 milliseconds
1100 MB mark took 7197 milliseconds
1200 MB mark took 7307 milliseconds
1300 MB mark took 7381 milliseconds
0 to 1300 MB took 95104 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 1435324409, requests: 40040, timestamps:
100 MB mark took 14006 milliseconds
200 MB mark took 13702 milliseconds
300 MB mark took 15301 milliseconds
400 MB mark took 14796 milliseconds
500 MB mark took 13901 milliseconds
600 MB mark took 13597 milliseconds
700 MB mark took 13589 milliseconds
800 MB mark took 13417 milliseconds
900 MB mark took 13691 milliseconds
1000 MB mark took 13605 milliseconds
1100 MB mark took 13487 milliseconds
1200 MB mark took 13502 milliseconds
1300 MB mark took 13410 milliseconds
0 to 1300 MB took 180007 milliseconds
```
</details>

#### bulk size = 1000

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 1443331150, requests: 2930, timestamps:
100 MB mark took 5779 milliseconds
200 MB mark took 5802 milliseconds
300 MB mark took 5802 milliseconds
400 MB mark took 5896 milliseconds
500 MB mark took 5807 milliseconds
600 MB mark took 5694 milliseconds
700 MB mark took 5708 milliseconds
800 MB mark took 5601 milliseconds
900 MB mark took 5813 milliseconds
1000 MB mark took 5873 milliseconds
1100 MB mark took 5503 milliseconds
1200 MB mark took 5903 milliseconds
1300 MB mark took 5700 milliseconds
0 to 1300 MB took 74886 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 1435324409, requests: 2930, timestamps:
100 MB mark took 11398 milliseconds
200 MB mark took 11520 milliseconds
300 MB mark took 11682 milliseconds
400 MB mark took 11005 milliseconds
500 MB mark took 10807 milliseconds
600 MB mark took 10995 milliseconds
700 MB mark took 10898 milliseconds
800 MB mark took 11199 milliseconds
900 MB mark took 10897 milliseconds
1000 MB mark took 10604 milliseconds
1100 MB mark took 11094 milliseconds
1200 MB mark took 10805 milliseconds
1300 MB mark took 10602 milliseconds
0 to 1300 MB took 143511 milliseconds
```
</details>

### 500 files 1 MB each

#### bulk = 50 (default)

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 1441779587, requests: 40060, timestamps:
100 MB mark took 14075 milliseconds
200 MB mark took 14391 milliseconds
300 MB mark took 12609 milliseconds
400 MB mark took 12995 milliseconds
500 MB mark took 14111 milliseconds
600 MB mark took 13899 milliseconds
700 MB mark took 13282 milliseconds
800 MB mark took 13317 milliseconds
900 MB mark took 13005 milliseconds
1000 MB mark took 12995 milliseconds
1100 MB mark took 12401 milliseconds
1200 MB mark took 13093 milliseconds
1300 MB mark took 18792 milliseconds
0 to 1300 MB took 178971 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 1433742399, requests: 40060, timestamps:
100 MB mark took 28078 milliseconds
200 MB mark took 29892 milliseconds
300 MB mark took 29103 milliseconds
400 MB mark took 27108 milliseconds
500 MB mark took 26394 milliseconds
600 MB mark took 28997 milliseconds
700 MB mark took 27097 milliseconds
800 MB mark took 27498 milliseconds
900 MB mark took 27405 milliseconds
1000 MB mark took 27400 milliseconds
1100 MB mark took 28302 milliseconds
1200 MB mark took 26997 milliseconds
1300 MB mark took 28606 milliseconds
0 to 1300 MB took 362884 milliseconds
```
</details>

#### bulk size = 1000

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 1441779587, requests: 2931, timestamps:
100 MB mark took 7900 milliseconds
200 MB mark took 9384 milliseconds
300 MB mark took 9494 milliseconds
400 MB mark took 9208 milliseconds
500 MB mark took 8807 milliseconds
600 MB mark took 9092 milliseconds
700 MB mark took 9097 milliseconds
800 MB mark took 8892 milliseconds
900 MB mark took 9710 milliseconds
1000 MB mark took 9095 milliseconds
1100 MB mark took 9101 milliseconds
1200 MB mark took 10905 milliseconds
1300 MB mark took 11903 milliseconds
0 to 1300 MB took 122594 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 1433742399, requests: 2931, timestamps:
100 MB mark took 19472 milliseconds
200 MB mark took 24792 milliseconds
300 MB mark took 25200 milliseconds
400 MB mark took 21519 milliseconds
500 MB mark took 22485 milliseconds
600 MB mark took 22099 milliseconds
700 MB mark took 22097 milliseconds
800 MB mark took 19907 milliseconds
900 MB mark took 23089 milliseconds
1000 MB mark took 22511 milliseconds
1100 MB mark took 21589 milliseconds
1200 MB mark took 23099 milliseconds
1300 MB mark took 21808 milliseconds
0 to 1300 MB took 289674 milliseconds
```
</details>
