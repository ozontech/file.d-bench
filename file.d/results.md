# File.d

## Bench 1 - Version TBD

### 1 file 500 MB

* Macbook Pro M1 cpu not limited
* Docker container CPU limit: none
* Docker container memory limit: 256 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 695759635, requests: 8298, timestamps:
100 MB mark took 47147 milliseconds
200 MB mark took 47383 milliseconds
300 MB mark took 47811 milliseconds
400 MB mark took 48562 milliseconds
500 MB mark took 50788 milliseconds
600 MB mark took 50273 milliseconds
0 to 600 MB took 291966 milliseconds
```
</details>

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 1.57%	Mem 3.004MiB / 512MiB = 0.59%	NetIO 5.67MB / 202kB	BlockIO 0B / 0B
filed-benchmark		CPU 108.92%	Mem 71.37MiB / 256MiB = 27.88%	NetIO 109kB / 2.91MB	BlockIO 23.2MB / 4.1kB
mockelasticsearch		CPU 1.73%	Mem 3.012MiB / 512MiB = 0.59%	NetIO 11.4MB / 405kB	BlockIO 0B / 0B
filed-benchmark		CPU 110.23%	Mem 74.59MiB / 256MiB = 29.14%	NetIO 495kB / 13.8MB	BlockIO 23.2MB / 24.6kB
mockelasticsearch		CPU 1.45%	Mem 3.027MiB / 512MiB = 0.59%	NetIO 23.1MB / 814kB	BlockIO 0B / 0B
filed-benchmark		CPU 112.93%	Mem 75.85MiB / 256MiB = 29.63%	NetIO 735kB / 20.9MB	BlockIO 23.2MB / 32.8kB
mockelasticsearch		CPU 1.20%	Mem 3.023MiB / 512MiB = 0.59%	NetIO 29.8MB / 1.05MB	BlockIO 0B / 0B
filed-benchmark		CPU 106.95%	Mem 77.54MiB / 256MiB = 30.29%	NetIO 1.12MB / 31.5MB	BlockIO 23.2MB / 53.2kB
mockelasticsearch		CPU 0.95%	Mem 3.023MiB / 512MiB = 0.59%	NetIO 39.9MB / 1.41MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.88%	Mem 78.72MiB / 256MiB = 30.75%	NetIO 1.36MB / 38.4MB	BlockIO 23.2MB / 65.5kB
mockelasticsearch		CPU 1.79%	Mem 3.117MiB / 512MiB = 0.61%	NetIO 48.8MB / 1.71MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.85%	Mem 79.75MiB / 256MiB = 31.15%	NetIO 1.62MB / 46.2MB	BlockIO 23.2MB / 77.8kB
mockelasticsearch		CPU 1.29%	Mem 3.059MiB / 512MiB = 0.60%	NetIO 56.4MB / 1.98MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.28%	Mem 80.93MiB / 256MiB = 31.61%	NetIO 1.9MB / 54MB	BlockIO 23.2MB / 94.2kB
mockelasticsearch		CPU 1.64%	Mem 3.035MiB / 512MiB = 0.59%	NetIO 62.3MB / 2.19MB	BlockIO 0B / 0B
filed-benchmark		CPU 114.54%	Mem 81.49MiB / 256MiB = 31.83%	NetIO 2.29MB / 65MB	BlockIO 23.2MB / 115kB
mockelasticsearch		CPU 1.26%	Mem 3.035MiB / 512MiB = 0.59%	NetIO 72.4MB / 2.56MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.86%	Mem 81.81MiB / 256MiB = 31.96%	NetIO 2.49MB / 70.8MB	BlockIO 23.2MB / 123kB
mockelasticsearch		CPU 1.18%	Mem 3.164MiB / 512MiB = 0.62%	NetIO 80.4MB / 2.84MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.09%	Mem 81.8MiB / 256MiB = 31.95%	NetIO 2.76MB / 78.1MB	BlockIO 23.2MB / 139kB
mockelasticsearch		CPU 1.10%	Mem 3.098MiB / 512MiB = 0.61%	NetIO 89.1MB / 3.15MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.89%	Mem 82.03MiB / 256MiB = 32.04%	NetIO 3.09MB / 87.2MB	BlockIO 23.2MB / 156kB
mockelasticsearch		CPU 1.25%	Mem 3.105MiB / 512MiB = 0.61%	NetIO 97.4MB / 3.43MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.75%	Mem 82.29MiB / 256MiB = 32.14%	NetIO 3.36MB / 95.2MB	BlockIO 23.2MB / 168kB
mockelasticsearch		CPU 1.39%	Mem 3.105MiB / 512MiB = 0.61%	NetIO 103MB / 3.63MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.26%	Mem 83.42MiB / 256MiB = 32.59%	NetIO 3.71MB / 105MB	BlockIO 23.2MB / 188kB
mockelasticsearch		CPU 1.51%	Mem 3.105MiB / 512MiB = 0.61%	NetIO 112MB / 3.94MB	BlockIO 0B / 0B
filed-benchmark		CPU 113.06%	Mem 84.7MiB / 256MiB = 33.08%	NetIO 4.05MB / 114MB	BlockIO 23.2MB / 205kB
mockelasticsearch		CPU 1.01%	Mem 3.105MiB / 512MiB = 0.61%	NetIO 122MB / 4.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.35%	Mem 85.48MiB / 256MiB = 33.39%	NetIO 4.23MB / 120MB	BlockIO 23.2MB / 213kB
mockelasticsearch		CPU 1.19%	Mem 3.105MiB / 512MiB = 0.61%	NetIO 131MB / 4.61MB	BlockIO 0B / 0B
filed-benchmark		CPU 113.29%	Mem 86.96MiB / 256MiB = 33.97%	NetIO 4.54MB / 128MB	BlockIO 23.2MB / 229kB
mockelasticsearch		CPU 1.84%	Mem 3.109MiB / 512MiB = 0.61%	NetIO 139MB / 4.89MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.33%	Mem 88.24MiB / 256MiB = 34.47%	NetIO 4.82MB / 137MB	BlockIO 23.2MB / 246kB
mockelasticsearch		CPU 3.20%	Mem 3.121MiB / 512MiB = 0.61%	NetIO 148MB / 5.19MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.27%	Mem 89.5MiB / 256MiB = 34.96%	NetIO 5.11MB / 145MB	BlockIO 23.2MB / 258kB
mockelasticsearch		CPU 2.30%	Mem 3.125MiB / 512MiB = 0.61%	NetIO 156MB / 5.48MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.07%	Mem 90.88MiB / 256MiB = 35.50%	NetIO 5.39MB / 153MB	BlockIO 23.2MB / 274kB
mockelasticsearch		CPU 1.18%	Mem 3.219MiB / 512MiB = 0.63%	NetIO 163MB / 5.75MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.95%	Mem 92.29MiB / 256MiB = 36.05%	NetIO 5.68MB / 161MB	BlockIO 23.2MB / 291kB
mockelasticsearch		CPU 1.14%	Mem 3.164MiB / 512MiB = 0.62%	NetIO 172MB / 6.05MB	BlockIO 0B / 0B
filed-benchmark		CPU 112.95%	Mem 93.61MiB / 256MiB = 36.56%	NetIO 5.99MB / 170MB	BlockIO 23.2MB / 303kB
mockelasticsearch		CPU 1.53%	Mem 3.129MiB / 512MiB = 0.61%	NetIO 178MB / 6.25MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.83%	Mem 95MiB / 256MiB = 37.11%	NetIO 6.25MB / 178MB	BlockIO 23.2MB / 319kB
mockelasticsearch		CPU 1.24%	Mem 3.125MiB / 512MiB = 0.61%	NetIO 187MB / 6.55MB	BlockIO 0B / 0B
filed-benchmark		CPU 117.22%	Mem 95.98MiB / 256MiB = 37.49%	NetIO 6.49MB / 185MB	BlockIO 23.2MB / 332kB
mockelasticsearch		CPU 1.43%	Mem 3.125MiB / 512MiB = 0.61%	NetIO 196MB / 6.87MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.95%	Mem 97.45MiB / 256MiB = 38.07%	NetIO 6.79MB / 194MB	BlockIO 23.2MB / 344kB
mockelasticsearch		CPU 1.59%	Mem 3.141MiB / 512MiB = 0.61%	NetIO 204MB / 7.16MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.09%	Mem 98.55MiB / 256MiB = 38.50%	NetIO 7.09MB / 202MB	BlockIO 23.2MB / 360kB
mockelasticsearch		CPU 1.51%	Mem 3.168MiB / 512MiB = 0.62%	NetIO 211MB / 7.42MB	BlockIO 0B / 0B
filed-benchmark		CPU 112.50%	Mem 99.89MiB / 256MiB = 39.02%	NetIO 7.33MB / 209MB	BlockIO 23.2MB / 377kB
mockelasticsearch		CPU 1.54%	Mem 3.152MiB / 512MiB = 0.62%	NetIO 220MB / 7.71MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.36%	Mem 100.8MiB / 256MiB = 39.39%	NetIO 7.62MB / 217MB	BlockIO 23.2MB / 389kB
mockelasticsearch		CPU 1.97%	Mem 3.156MiB / 512MiB = 0.62%	NetIO 229MB / 8.02MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.48%	Mem 102.5MiB / 256MiB = 40.03%	NetIO 7.92MB / 226MB	BlockIO 23.2MB / 406kB
mockelasticsearch		CPU 1.15%	Mem 3.168MiB / 512MiB = 0.62%	NetIO 237MB / 8.32MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.60%	Mem 104.1MiB / 256MiB = 40.65%	NetIO 8.25MB / 236MB	BlockIO 23.2MB / 422kB
mockelasticsearch		CPU 2.04%	Mem 3.168MiB / 512MiB = 0.62%	NetIO 246MB / 8.64MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.03%	Mem 105.6MiB / 256MiB = 41.25%	NetIO 8.52MB / 243MB	BlockIO 23.2MB / 434kB
mockelasticsearch		CPU 1.67%	Mem 3.184MiB / 512MiB = 0.62%	NetIO 253MB / 8.86MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.83%	Mem 107MiB / 256MiB = 41.78%	NetIO 8.78MB / 251MB	BlockIO 23.2MB / 451kB
mockelasticsearch		CPU 1.86%	Mem 3.184MiB / 512MiB = 0.62%	NetIO 262MB / 9.18MB	BlockIO 0B / 0B
filed-benchmark		CPU 113.99%	Mem 108.4MiB / 256MiB = 42.34%	NetIO 9.08MB / 259MB	BlockIO 23.2MB / 467kB
mockelasticsearch		CPU 1.27%	Mem 3.191MiB / 512MiB = 0.62%	NetIO 272MB / 9.53MB	BlockIO 0B / 0B
filed-benchmark		CPU 133.62%	Mem 110.6MiB / 256MiB = 43.19%	NetIO 9.46MB / 270MB	BlockIO 23.2MB / 479kB
mockelasticsearch		CPU 1.11%	Mem 3.191MiB / 512MiB = 0.62%	NetIO 278MB / 9.72MB	BlockIO 0B / 0B
filed-benchmark		CPU 106.76%	Mem 111.3MiB / 256MiB = 43.50%	NetIO 9.78MB / 280MB	BlockIO 23.2MB / 500kB
mockelasticsearch		CPU 1.24%	Mem 3.215MiB / 512MiB = 0.63%	NetIO 285MB / 9.99MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.14%	Mem 111.8MiB / 256MiB = 43.66%	NetIO 10.1MB / 287MB	BlockIO 23.2MB / 516kB
mockelasticsearch		CPU 2.06%	Mem 3.191MiB / 512MiB = 0.62%	NetIO 296MB / 10.4MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.76%	Mem 112MiB / 256MiB = 43.75%	NetIO 10.3MB / 293MB	BlockIO 23.2MB / 524kB
mockelasticsearch		CPU 1.45%	Mem 3.203MiB / 512MiB = 0.63%	NetIO 303MB / 10.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 120.62%	Mem 112.1MiB / 256MiB = 43.81%	NetIO 10.5MB / 301MB	BlockIO 23.2MB / 541kB
mockelasticsearch		CPU 1.02%	Mem 3.215MiB / 512MiB = 0.63%	NetIO 310MB / 10.8MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.92%	Mem 112.4MiB / 256MiB = 43.90%	NetIO 10.8MB / 308MB	BlockIO 23.2MB / 557kB
mockelasticsearch		CPU 1.79%	Mem 3.219MiB / 512MiB = 0.63%	NetIO 318MB / 11.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.35%	Mem 112.4MiB / 256MiB = 43.91%	NetIO 11.1MB / 316MB	BlockIO 23.2MB / 569kB
mockelasticsearch		CPU 2.11%	Mem 3.27MiB / 512MiB = 0.64%	NetIO 329MB / 11.5MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.54%	Mem 112.8MiB / 256MiB = 44.05%	NetIO 11.4MB / 325MB	BlockIO 23.2MB / 586kB
mockelasticsearch		CPU 1.86%	Mem 3.227MiB / 512MiB = 0.63%	NetIO 337MB / 11.8MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.60%	Mem 112.9MiB / 256MiB = 44.08%	NetIO 11.7MB / 334MB	BlockIO 23.2MB / 602kB
mockelasticsearch		CPU 1.03%	Mem 3.328MiB / 512MiB = 0.65%	NetIO 345MB / 12.1MB	BlockIO 0B / 0B
filed-benchmark		CPU 115.09%	Mem 113.3MiB / 256MiB = 44.24%	NetIO 12MB / 343MB	BlockIO 23.2MB / 618kB
mockelasticsearch		CPU 1.08%	Mem 3.23MiB / 512MiB = 0.63%	NetIO 353MB / 12.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.87%	Mem 113.4MiB / 256MiB = 44.29%	NetIO 12.3MB / 351MB	BlockIO 23.2MB / 631kB
mockelasticsearch		CPU 1.85%	Mem 3.355MiB / 512MiB = 0.66%	NetIO 362MB / 12.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.98%	Mem 113.6MiB / 256MiB = 44.38%	NetIO 12.5MB / 360MB	BlockIO 23.2MB / 647kB
mockelasticsearch		CPU 1.85%	Mem 3.238MiB / 512MiB = 0.63%	NetIO 370MB / 12.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 106.92%	Mem 113.7MiB / 256MiB = 44.43%	NetIO 12.8MB / 367MB	BlockIO 23.2MB / 664kB
mockelasticsearch		CPU 2.03%	Mem 3.273MiB / 512MiB = 0.64%	NetIO 378MB / 13.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.19%	Mem 114MiB / 256MiB = 44.53%	NetIO 13.1MB / 375MB	BlockIO 23.2MB / 676kB
mockelasticsearch		CPU 1.15%	Mem 3.242MiB / 512MiB = 0.63%	NetIO 386MB / 13.4MB	BlockIO 0B / 0B
filed-benchmark		CPU 112.18%	Mem 115.3MiB / 256MiB = 45.04%	NetIO 13.4MB / 384MB	BlockIO 23.2MB / 692kB
mockelasticsearch		CPU 1.29%	Mem 3.305MiB / 512MiB = 0.65%	NetIO 391MB / 13.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.02%	Mem 116.5MiB / 256MiB = 45.49%	NetIO 13.7MB / 393MB	BlockIO 23.5MB / 709kB
mockelasticsearch		CPU 1.69%	Mem 3.25MiB / 512MiB = 0.63%	NetIO 401MB / 13.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.76%	Mem 116.7MiB / 256MiB = 45.59%	NetIO 13.9MB / 398MB	BlockIO 23.5MB / 721kB
mockelasticsearch		CPU 1.67%	Mem 3.25MiB / 512MiB = 0.63%	NetIO 407MB / 14.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.64%	Mem 116.7MiB / 256MiB = 45.59%	NetIO 14.1MB / 405MB	BlockIO 23.5MB / 737kB
mockelasticsearch		CPU 1.22%	Mem 3.25MiB / 512MiB = 0.63%	NetIO 415MB / 14.5MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.34%	Mem 116.9MiB / 256MiB = 45.68%	NetIO 14.4MB / 413MB	BlockIO 23.5MB / 750kB
mockelasticsearch		CPU 1.51%	Mem 3.254MiB / 512MiB = 0.64%	NetIO 421MB / 14.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.95%	Mem 117.6MiB / 256MiB = 45.92%	NetIO 14.7MB / 423MB	BlockIO 23.6MB / 770kB
mockelasticsearch		CPU 1.39%	Mem 3.262MiB / 512MiB = 0.64%	NetIO 432MB / 15MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.47%	Mem 117.8MiB / 256MiB = 46.03%	NetIO 15MB / 430MB	BlockIO 23.6MB / 782kB
mockelasticsearch		CPU 1.75%	Mem 3.438MiB / 512MiB = 0.67%	NetIO 440MB / 15.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.03%	Mem 118MiB / 256MiB = 46.08%	NetIO 15.2MB / 438MB	BlockIO 23.6MB / 799kB
mockelasticsearch		CPU 1.36%	Mem 3.273MiB / 512MiB = 0.64%	NetIO 448MB / 15.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 134.28%	Mem 118.1MiB / 256MiB = 46.14%	NetIO 15.5MB / 446MB	BlockIO 23.8MB / 811kB
mockelasticsearch		CPU 1.75%	Mem 3.273MiB / 512MiB = 0.64%	NetIO 454MB / 15.8MB	BlockIO 0B / 0B
filed-benchmark		CPU 112.61%	Mem 118.7MiB / 256MiB = 46.38%	NetIO 15.8MB / 454MB	BlockIO 24MB / 827kB
mockelasticsearch		CPU 1.68%	Mem 3.281MiB / 512MiB = 0.64%	NetIO 462MB / 16.1MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.75%	Mem 118.9MiB / 256MiB = 46.46%	NetIO 16MB / 460MB	BlockIO 24MB / 840kB
mockelasticsearch		CPU 0.91%	Mem 3.289MiB / 512MiB = 0.64%	NetIO 469MB / 16.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.94%	Mem 119.1MiB / 256MiB = 46.52%	NetIO 16.3MB / 468MB	BlockIO 24.1MB / 852kB
mockelasticsearch		CPU 2.02%	Mem 3.289MiB / 512MiB = 0.64%	NetIO 478MB / 16.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.68%	Mem 119.3MiB / 256MiB = 46.61%	NetIO 16.6MB / 475MB	BlockIO 24.1MB / 868kB
mockelasticsearch		CPU 1.83%	Mem 3.293MiB / 512MiB = 0.64%	NetIO 486MB / 16.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.65%	Mem 119.5MiB / 256MiB = 46.68%	NetIO 16.8MB / 483MB	BlockIO 24.2MB / 885kB
mockelasticsearch		CPU 1.40%	Mem 3.289MiB / 512MiB = 0.64%	NetIO 494MB / 17.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.59%	Mem 119.7MiB / 256MiB = 46.75%	NetIO 17.1MB / 492MB	BlockIO 24.2MB / 897kB
mockelasticsearch		CPU 2.34%	Mem 3.305MiB / 512MiB = 0.65%	NetIO 501MB / 17.5MB	BlockIO 0B / 0B
filed-benchmark		CPU 115.75%	Mem 120.2MiB / 256MiB = 46.93%	NetIO 17.4MB / 499MB	BlockIO 24.3MB / 913kB
mockelasticsearch		CPU 1.03%	Mem 3.312MiB / 512MiB = 0.65%	NetIO 508MB / 17.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.50%	Mem 120.9MiB / 256MiB = 47.24%	NetIO 17.7MB / 507MB	BlockIO 24.4MB / 930kB
mockelasticsearch		CPU 1.13%	Mem 3.312MiB / 512MiB = 0.65%	NetIO 514MB / 17.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 106.95%	Mem 122.2MiB / 256MiB = 47.72%	NetIO 18MB / 516MB	BlockIO 24.6MB / 946kB
mockelasticsearch		CPU 1.37%	Mem 3.312MiB / 512MiB = 0.65%	NetIO 523MB / 18.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.00%	Mem 122.5MiB / 256MiB = 47.84%	NetIO 18.2MB / 521MB	BlockIO 24.6MB / 958kB
mockelasticsearch		CPU 1.51%	Mem 3.32MiB / 512MiB = 0.65%	NetIO 529MB / 18.4MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.18%	Mem 122.5MiB / 256MiB = 47.86%	NetIO 18.5MB / 531MB	BlockIO 24.6MB / 979kB
mockelasticsearch		CPU 0.97%	Mem 3.32MiB / 512MiB = 0.65%	NetIO 536MB / 18.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.73%	Mem 123MiB / 256MiB = 48.05%	NetIO 18.7MB / 538MB	BlockIO 24.6MB / 991kB
mockelasticsearch		CPU 1.24%	Mem 3.617MiB / 512MiB = 0.71%	NetIO 546MB / 18.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.42%	Mem 123.1MiB / 256MiB = 48.07%	NetIO 18.9MB / 544MB	BlockIO 24.6MB / 1MB
mockelasticsearch		CPU 1.31%	Mem 3.387MiB / 512MiB = 0.66%	NetIO 554MB / 19.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.16%	Mem 123.5MiB / 256MiB = 48.23%	NetIO 19.1MB / 552MB	BlockIO 24.6MB / 1.02MB
mockelasticsearch		CPU 2.40%	Mem 3.383MiB / 512MiB = 0.66%	NetIO 564MB / 19.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 116.01%	Mem 123.3MiB / 256MiB = 48.16%	NetIO 19.5MB / 560MB	BlockIO 24.6MB / 1.03MB
mockelasticsearch		CPU 1.15%	Mem 3.383MiB / 512MiB = 0.66%	NetIO 571MB / 19.8MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.69%	Mem 123.9MiB / 256MiB = 48.40%	NetIO 19.8MB / 569MB	BlockIO 24.6MB / 1.05MB
mockelasticsearch		CPU 1.35%	Mem 3.391MiB / 512MiB = 0.66%	NetIO 579MB / 20.1MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.34%	Mem 124.1MiB / 256MiB = 48.46%	NetIO 20.1MB / 577MB	BlockIO 24.7MB / 1.06MB
mockelasticsearch		CPU 1.30%	Mem 3.395MiB / 512MiB = 0.66%	NetIO 587MB / 20.4MB	BlockIO 0B / 0B
filed-benchmark		CPU 106.35%	Mem 124MiB / 256MiB = 48.43%	NetIO 20.3MB / 585MB	BlockIO 24.7MB / 1.08MB
mockelasticsearch		CPU 1.24%	Mem 3.395MiB / 512MiB = 0.66%	NetIO 594MB / 20.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.93%	Mem 124.5MiB / 256MiB = 48.62%	NetIO 20.6MB / 592MB	BlockIO 24.7MB / 1.09MB
mockelasticsearch		CPU 1.06%	Mem 3.398MiB / 512MiB = 0.66%	NetIO 602MB / 20.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.67%	Mem 124.4MiB / 256MiB = 48.60%	NetIO 20.9MB / 600MB	BlockIO 24.7MB / 1.11MB
mockelasticsearch		CPU 1.02%	Mem 3.43MiB / 512MiB = 0.67%	NetIO 609MB / 21.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.32%	Mem 124.9MiB / 256MiB = 48.77%	NetIO 21.1MB / 607MB	BlockIO 24.7MB / 1.12MB
mockelasticsearch		CPU 1.42%	Mem 3.398MiB / 512MiB = 0.66%	NetIO 617MB / 21.5MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.55%	Mem 124.9MiB / 256MiB = 48.78%	NetIO 21.4MB / 615MB	BlockIO 24.9MB / 1.14MB
mockelasticsearch		CPU 1.66%	Mem 3.418MiB / 512MiB = 0.67%	NetIO 625MB / 21.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.87%	Mem 125.1MiB / 256MiB = 48.86%	NetIO 21.7MB / 623MB	BlockIO 25MB / 1.16MB
mockelasticsearch		CPU 1.21%	Mem 3.418MiB / 512MiB = 0.67%	NetIO 632MB / 22MB	BlockIO 0B / 0B
filed-benchmark		CPU 109.65%	Mem 125.5MiB / 256MiB = 49.01%	NetIO 21.9MB / 630MB	BlockIO 25MB / 1.17MB
mockelasticsearch		CPU 0.88%	Mem 3.5MiB / 512MiB = 0.68%	NetIO 640MB / 22.2MB	BlockIO 0B / 0B
filed-benchmark		CPU 112.28%	Mem 125.8MiB / 256MiB = 49.14%	NetIO 22.2MB / 638MB	BlockIO 25.1MB / 1.18MB
mockelasticsearch		CPU 1.44%	Mem 3.422MiB / 512MiB = 0.67%	NetIO 645MB / 22.4MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.39%	Mem 126MiB / 256MiB = 49.20%	NetIO 22.5MB / 647MB	BlockIO 25.1MB / 1.2MB
mockelasticsearch		CPU 1.32%	Mem 3.426MiB / 512MiB = 0.67%	NetIO 655MB / 22.8MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.52%	Mem 125.9MiB / 256MiB = 49.16%	NetIO 22.7MB / 653MB	BlockIO 25.1MB / 1.21MB
mockelasticsearch		CPU 1.02%	Mem 3.43MiB / 512MiB = 0.67%	NetIO 661MB / 23MB	BlockIO 0B / 0B
filed-benchmark		CPU 110.92%	Mem 126.3MiB / 256MiB = 49.32%	NetIO 23.1MB / 663MB	BlockIO 25.1MB / 1.23MB
mockelasticsearch		CPU 1.86%	Mem 3.43MiB / 512MiB = 0.67%	NetIO 671MB / 23.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.25%	Mem 126.3MiB / 256MiB = 49.35%	NetIO 23.2MB / 668MB	BlockIO 25.1MB / 1.25MB
mockelasticsearch		CPU 1.26%	Mem 3.434MiB / 512MiB = 0.67%	NetIO 679MB / 23.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.54%	Mem 126.5MiB / 256MiB = 49.41%	NetIO 23.5MB / 677MB	BlockIO 25.1MB / 1.26MB
mockelasticsearch		CPU 1.60%	Mem 3.453MiB / 512MiB = 0.67%	NetIO 687MB / 23.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.94%	Mem 126.9MiB / 256MiB = 49.57%	NetIO 23.8MB / 685MB	BlockIO 25.1MB / 1.27MB
mockelasticsearch		CPU 1.04%	Mem 3.453MiB / 512MiB = 0.67%	NetIO 692MB / 24.1MB	BlockIO 0B / 0B
filed-benchmark		CPU 111.39%	Mem 127.1MiB / 256MiB = 49.64%	NetIO 24.1MB / 694MB	BlockIO 25.1MB / 1.29MB
mockelasticsearch		CPU 1.44%	Mem 3.453MiB / 512MiB = 0.67%	NetIO 701MB / 24.4MB	BlockIO 0B / 0B
filed-benchmark		CPU 107.13%	Mem 127.1MiB / 256MiB = 49.65%	NetIO 24.3MB / 699MB	BlockIO 25.1MB / 1.3MB
mockelasticsearch		CPU 0.96%	Mem 3.457MiB / 512MiB = 0.68%	NetIO 708MB / 24.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 108.42%	Mem 127.4MiB / 256MiB = 49.78%	NetIO 24.6MB / 707MB	BlockIO 25.1MB / 1.32MB
mockelasticsearch		CPU 0.10%	Mem 3.461MiB / 512MiB = 0.68%	NetIO 709MB / 24.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 1.32%	Mem 127.4MiB / 256MiB = 49.77%	NetIO 24.7MB / 709MB	BlockIO 25.2MB / 1.33MB
```
</details>

### 500 files 1 MB each

* Macbook Pro M1 cpu not limited
* Docker container CPU limit: none
* Docker container memory limit: 256 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 687983117, requests: 7818, timestamps:
100 MB mark took 8427 milliseconds
200 MB mark took 6996 milliseconds
300 MB mark took 7085 milliseconds
400 MB mark took 6898 milliseconds
500 MB mark took 7039 milliseconds
600 MB mark took 6768 milliseconds
0 to 600 MB took 43216 milliseconds
```
</details>

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 0.71%	Mem 1.855MiB / 512MiB = 0.36%	NetIO 683kB / 27.4kB	BlockIO 0B / 0B
filed-benchmark		CPU 261.74%	Mem 73.42MiB / 256MiB = 28.68%	NetIO 429kB / 12.7MB	BlockIO 102kB / 8.19kB
mockelasticsearch		CPU 8.92%	Mem 2.879MiB / 512MiB = 0.56%	NetIO 60.5MB / 2MB	BlockIO 0B / 0B
filed-benchmark		CPU 285.40%	Mem 78.51MiB / 256MiB = 30.67%	NetIO 1.57MB / 47.1MB	BlockIO 102kB / 45.1kB
mockelasticsearch		CPU 8.28%	Mem 2.883MiB / 512MiB = 0.56%	NetIO 113MB / 3.76MB	BlockIO 0B / 0B
filed-benchmark		CPU 298.42%	Mem 80.71MiB / 256MiB = 31.53%	NetIO 3.28MB / 98.6MB	BlockIO 102kB / 213kB
mockelasticsearch		CPU 8.51%	Mem 3.031MiB / 512MiB = 0.59%	NetIO 171MB / 5.63MB	BlockIO 0B / 0B
filed-benchmark		CPU 310.69%	Mem 89.08MiB / 256MiB = 34.80%	NetIO 5.13MB / 155MB	BlockIO 102kB / 426kB
mockelasticsearch		CPU 8.68%	Mem 3.031MiB / 512MiB = 0.59%	NetIO 211MB / 6.87MB	BlockIO 0B / 0B
filed-benchmark		CPU 305.98%	Mem 99.29MiB / 256MiB = 38.78%	NetIO 7.34MB / 226MB	BlockIO 164kB / 639kB
mockelasticsearch		CPU 8.88%	Mem 3.059MiB / 512MiB = 0.60%	NetIO 282MB / 9.1MB	BlockIO 0B / 0B
filed-benchmark		CPU 306.18%	Mem 105.8MiB / 256MiB = 41.33%	NetIO 8.62MB / 267MB	BlockIO 164kB / 799kB
mockelasticsearch		CPU 8.70%	Mem 3.047MiB / 512MiB = 0.60%	NetIO 323MB / 10.3MB	BlockIO 0B / 0B
filed-benchmark		CPU 309.28%	Mem 109.8MiB / 256MiB = 42.90%	NetIO 10.8MB / 339MB	BlockIO 164kB / 1.06MB
mockelasticsearch		CPU 9.19%	Mem 3.195MiB / 512MiB = 0.62%	NetIO 395MB / 12.7MB	BlockIO 0B / 0B
filed-benchmark		CPU 311.28%	Mem 110.7MiB / 256MiB = 43.24%	NetIO 12.2MB / 380MB	BlockIO 164kB / 1.17MB
mockelasticsearch		CPU 9.36%	Mem 3.148MiB / 512MiB = 0.61%	NetIO 453MB / 14.5MB	BlockIO 0B / 0B
filed-benchmark		CPU 306.02%	Mem 112.1MiB / 256MiB = 43.77%	NetIO 14MB / 437MB	BlockIO 164kB / 1.38MB
mockelasticsearch		CPU 9.51%	Mem 3.16MiB / 512MiB = 0.62%	NetIO 494MB / 15.8MB	BlockIO 0B / 0B
filed-benchmark		CPU 303.27%	Mem 114MiB / 256MiB = 44.53%	NetIO 16.3MB / 509MB	BlockIO 164kB / 1.65MB
mockelasticsearch		CPU 10.09%	Mem 3.273MiB / 512MiB = 0.64%	NetIO 549MB / 17.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 312.54%	Mem 119.6MiB / 256MiB = 46.71%	NetIO 18.1MB / 565MB	BlockIO 426kB / 1.81MB
mockelasticsearch		CPU 9.51%	Mem 3.273MiB / 512MiB = 0.64%	NetIO 623MB / 19.9MB	BlockIO 0B / 0B
filed-benchmark		CPU 311.95%	Mem 120.5MiB / 256MiB = 47.08%	NetIO 19.4MB / 606MB	BlockIO 426kB / 1.97MB
mockelasticsearch		CPU 9.53%	Mem 3.266MiB / 512MiB = 0.64%	NetIO 685MB / 22.1MB	BlockIO 0B / 0B
filed-benchmark		CPU 374.14%	Mem 121.1MiB / 256MiB = 47.31%	NetIO 21.5MB / 668MB	BlockIO 426kB / 2.18MB
mockelasticsearch		CPU 0.13%	Mem 3.258MiB / 512MiB = 0.64%	NetIO 701MB / 22.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 1.24%	Mem 122.4MiB / 256MiB = 47.81%	NetIO 22.6MB / 701MB	BlockIO 426kB / 2.34MB
mockelasticsearch		CPU 0.05%	Mem 3.266MiB / 512MiB = 0.64%	NetIO 701MB / 22.6MB	BlockIO 0B / 0B
filed-benchmark		CPU 8.35%	Mem 122.4MiB / 256MiB = 47.82%	NetIO 22.6MB / 701MB	BlockIO 426kB / 2.34MB
```
</details>


## Bench 2 - Version TBD

### 1 file 500 MB

#### Macbook Pro M1

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 687759601, requests: 8270, timestamps:
100 MB mark took 51892 milliseconds
200 MB mark took 48415 milliseconds
300 MB mark took 46906 milliseconds
400 MB mark took 48583 milliseconds
500 MB mark took 47114 milliseconds
600 MB mark took 50207 milliseconds
0 to 600 MB took 293119 milliseconds
```
</details>

#### Intel Xeon Processor

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 679759595, requests: 8779, timestamps:
100 MB mark took 72769 milliseconds
200 MB mark took 88741 milliseconds
300 MB mark took 102153 milliseconds
400 MB mark took 102100 milliseconds
500 MB mark took 101616 milliseconds
600 MB mark took 65488 milliseconds
0 to 600 MB took 532870 milliseconds
```
</details>

### 500 files 1 MB each

#### Macbook Pro M1

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 687982803, requests: 7948, timestamps:
100 MB mark took 54194 milliseconds
200 MB mark took 56904 milliseconds
300 MB mark took 56791 milliseconds
400 MB mark took 58198 milliseconds
500 MB mark took 59601 milliseconds
600 MB mark took 69605 milliseconds
0 to 600 MB took 355295 milliseconds
```
</details>

#### Intel Xeon Processor

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 679978679, requests: 7817, timestamps:
100 MB mark took 12572 milliseconds
200 MB mark took 12286 milliseconds
300 MB mark took 12208 milliseconds
400 MB mark took 12379 milliseconds
500 MB mark took 12297 milliseconds
600 MB mark took 11946 milliseconds
0 to 600 MB took 73691 milliseconds
```
</details>
