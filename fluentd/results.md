# Fluentd

## Bench 1 - Version 1.14-1

### single file 500 MB

```
STATS = overall bytes: 701101032, requests: 207, timestamps:
100 MB mark took 64716 milliseconds
200 MB mark took 66717 milliseconds
300 MB mark took 66474 milliseconds
400 MB mark took 67855 milliseconds
500 MB mark took 71716 milliseconds
600 MB mark took 63288 milliseconds
0 to 600 MB took 400769 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 0.00%	Mem 1.516MiB / 512MiB = 0.30%	NetIO 1.04kB / 0B	BlockIO 61.4kB / 0B
fluentd-benchmark		CPU 99.59%	Mem 63.96MiB / 256MiB = 24.99%	NetIO 1.04kB / 0B	BlockIO 0B / 0B
mockelasticsearch		CPU 0.00%	Mem 1.516MiB / 512MiB = 0.30%	NetIO 1.11kB / 0B	BlockIO 61.4kB / 0B
fluentd-benchmark		CPU 99.89%	Mem 125MiB / 256MiB = 48.83%	NetIO 1.11kB / 0B	BlockIO 0B / 0B
mockelasticsearch		CPU 0.05%	Mem 1.656MiB / 512MiB = 0.32%	NetIO 2.33kB / 1.78kB	BlockIO 61.4kB / 0B
fluentd-benchmark		CPU 103.20%	Mem 155.7MiB / 256MiB = 60.82%	NetIO 2.94kB / 1.22kB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.36%	Mem 6.449MiB / 512MiB = 1.26%	NetIO 7.8MB / 17kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.20%	Mem 155.6MiB / 256MiB = 60.79%	NetIO 11.1kB / 3.9MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.03%	Mem 10.62MiB / 512MiB = 2.07%	NetIO 10.2MB / 23.9kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.52%	Mem 166.2MiB / 256MiB = 64.90%	NetIO 25.2kB / 10.2MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.23%	Mem 10.93MiB / 512MiB = 2.14%	NetIO 20.4MB / 43.9kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.33%	Mem 163MiB / 256MiB = 63.67%	NetIO 38.7kB / 18MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.01%	Mem 10.93MiB / 512MiB = 2.14%	NetIO 20.4MB / 44kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.68%	Mem 170MiB / 256MiB = 66.42%	NetIO 51.5kB / 24.3MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.20%	Mem 10.93MiB / 512MiB = 2.14%	NetIO 30.6MB / 63.3kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.94%	Mem 159MiB / 256MiB = 62.11%	NetIO 57.9kB / 28.2MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.01%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 34.4MB / 69.6kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.48%	Mem 176MiB / 256MiB = 68.74%	NetIO 70.9kB / 34.4MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.04%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 40.7MB / 82.2kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.31%	Mem 175.1MiB / 256MiB = 68.41%	NetIO 83.5kB / 40.7MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.32%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 48.5MB / 95.6kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.25%	Mem 166MiB / 256MiB = 64.86%	NetIO 90.7kB / 44.6MB	BlockIO 0B / 0B
mockelasticsearch		CPU 0.01%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 50.9MB / 101kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.88%	Mem 174.9MiB / 256MiB = 68.32%	NetIO 102kB / 50.9MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.01%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 58.7MB / 114kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.58%	Mem 167.6MiB / 256MiB = 65.47%	NetIO 116kB / 58.7MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.33%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 65MB / 128kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.62%	Mem 171.1MiB / 256MiB = 66.84%	NetIO 122kB / 61.1MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.24%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 71.3MB / 141kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.45%	Mem 172.6MiB / 256MiB = 67.42%	NetIO 136kB / 68.9MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.02%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 75.2MB / 148kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.26%	Mem 180.5MiB / 256MiB = 70.52%	NetIO 150kB / 75.2MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.25%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 81.5MB / 161kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.93%	Mem 172.3MiB / 256MiB = 67.31%	NetIO 162kB / 81.5MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.28%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 89.3MB / 175kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.82%	Mem 174.7MiB / 256MiB = 68.23%	NetIO 169kB / 85.4MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.01%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 91.6MB / 182kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.43%	Mem 165.4MiB / 256MiB = 64.61%	NetIO 183kB / 91.6MB	BlockIO 0B / 4.1kB
mockelasticsearch		CPU 0.24%	Mem 10.94MiB / 512MiB = 2.14%	NetIO 99.4MB / 195kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.42%	Mem 167.9MiB / 256MiB = 65.57%	NetIO 189kB / 95.5MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.30%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 106MB / 208kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.79%	Mem 176.3MiB / 256MiB = 68.86%	NetIO 203kB / 102MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.23%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 112MB / 221kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.60%	Mem 177.4MiB / 256MiB = 69.30%	NetIO 216kB / 110MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.27%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 116MB / 227kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.66%	Mem 167.9MiB / 256MiB = 65.60%	NetIO 222kB / 112MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.02%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 122MB / 241kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.90%	Mem 171.6MiB / 256MiB = 67.03%	NetIO 242kB / 122MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.02%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 126MB / 248kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.83%	Mem 178.4MiB / 256MiB = 69.69%	NetIO 249kB / 126MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.07%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 132MB / 262kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.02%	Mem 175.5MiB / 256MiB = 68.56%	NetIO 263kB / 132MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.31%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 140MB / 274kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.92%	Mem 178.5MiB / 256MiB = 69.73%	NetIO 270kB / 136MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.26%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 146MB / 287kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.99%	Mem 186.1MiB / 256MiB = 72.69%	NetIO 281kB / 143MB	BlockIO 0B / 8.19kB
mockelasticsearch		CPU 0.21%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 153MB / 299kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.81%	Mem 166.2MiB / 256MiB = 64.94%	NetIO 295kB / 150MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.30%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 157MB / 305kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.20%	Mem 178.7MiB / 256MiB = 69.82%	NetIO 300kB / 153MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.22%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 163MB / 318kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.95%	Mem 186.1MiB / 256MiB = 72.68%	NetIO 319kB / 163MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.01%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 167MB / 324kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.16%	Mem 166.9MiB / 256MiB = 65.20%	NetIO 326kB / 167MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.02%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 173MB / 337kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 103.40%	Mem 176.8MiB / 256MiB = 69.07%	NetIO 339kB / 173MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.30%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 181MB / 351kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.11%	Mem 181MiB / 256MiB = 70.69%	NetIO 346kB / 177MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.01%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 183MB / 358kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.49%	Mem 168MiB / 256MiB = 65.62%	NetIO 359kB / 183MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.34%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 191MB / 371kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.41%	Mem 177.9MiB / 256MiB = 69.50%	NetIO 378kB / 193MB	BlockIO 0B / 12.3kB
mockelasticsearch		CPU 0.32%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 197MB / 383kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.85%	Mem 181.6MiB / 256MiB = 70.95%	NetIO 378kB / 193MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.03%	Mem 10.95MiB / 512MiB = 2.14%	NetIO 201MB / 389kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.94%	Mem 189.7MiB / 256MiB = 74.09%	NetIO 396kB / 204MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.01%	Mem 11.05MiB / 512MiB = 2.16%	NetIO 208MB / 402kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.35%	Mem 175.6MiB / 256MiB = 68.61%	NetIO 403kB / 208MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.02%	Mem 11.06MiB / 512MiB = 2.16%	NetIO 214MB / 415kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.11%	Mem 176.6MiB / 256MiB = 68.99%	NetIO 416kB / 214MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.31%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 222MB / 429kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.02%	Mem 180.9MiB / 256MiB = 70.68%	NetIO 423kB / 218MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.01%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 224MB / 436kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.32%	Mem 189.2MiB / 256MiB = 73.91%	NetIO 437kB / 224MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.28%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 232MB / 450kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.29%	Mem 172.6MiB / 256MiB = 67.40%	NetIO 445kB / 228MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.21%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 237MB / 462kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.78%	Mem 181MiB / 256MiB = 70.71%	NetIO 459kB / 234MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.27%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 244MB / 479kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.96%	Mem 182.3MiB / 256MiB = 71.20%	NetIO 473kB / 242MB	BlockIO 4.1kB / 16.4kB
mockelasticsearch		CPU 0.24%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 248MB / 484kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.16%	Mem 174.8MiB / 256MiB = 68.27%	NetIO 480kB / 244MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.25%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 255MB / 497kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.64%	Mem 177MiB / 256MiB = 69.14%	NetIO 498kB / 255MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.28%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 258MB / 504kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.77%	Mem 184.8MiB / 256MiB = 72.17%	NetIO 498kB / 255MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.02%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 265MB / 518kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.80%	Mem 186.1MiB / 256MiB = 72.71%	NetIO 520kB / 265MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.01%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 269MB / 525kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.62%	Mem 182.1MiB / 256MiB = 71.13%	NetIO 527kB / 269MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.01%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 275MB / 536kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.09%	Mem 182.2MiB / 256MiB = 71.19%	NetIO 538kB / 275MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.32%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 283MB / 551kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.76%	Mem 187.2MiB / 256MiB = 73.11%	NetIO 545kB / 279MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.01%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 285MB / 557kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.13%	Mem 175MiB / 256MiB = 68.36%	NetIO 558kB / 285MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.01%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 293MB / 568kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.14%	Mem 185MiB / 256MiB = 72.26%	NetIO 570kB / 293MB	BlockIO 4.1kB / 20.5kB
mockelasticsearch		CPU 0.37%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 299MB / 581kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.76%	Mem 187.5MiB / 256MiB = 73.23%	NetIO 575kB / 295MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.22%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 305MB / 594kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.23%	Mem 188.8MiB / 256MiB = 73.75%	NetIO 588kB / 303MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.02%	Mem 15.24MiB / 512MiB = 2.98%	NetIO 309MB / 601kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.03%	Mem 183.5MiB / 256MiB = 71.69%	NetIO 602kB / 309MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.01%	Mem 15.29MiB / 512MiB = 2.99%	NetIO 316MB / 614kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.42%	Mem 183.3MiB / 256MiB = 71.60%	NetIO 616kB / 316MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.30%	Mem 15.29MiB / 512MiB = 2.99%	NetIO 323MB / 629kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.54%	Mem 189.3MiB / 256MiB = 73.96%	NetIO 623kB / 320MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.03%	Mem 15.29MiB / 512MiB = 2.99%	NetIO 326MB / 635kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.16%	Mem 179.1MiB / 256MiB = 69.95%	NetIO 637kB / 326MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.26%	Mem 19.57MiB / 512MiB = 3.82%	NetIO 334MB / 650kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.38%	Mem 186.2MiB / 256MiB = 72.72%	NetIO 644kB / 330MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.31%	Mem 19.57MiB / 512MiB = 3.82%	NetIO 340MB / 663kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.68%	Mem 193.7MiB / 256MiB = 75.67%	NetIO 657kB / 336MB	BlockIO 4.1kB / 24.6kB
mockelasticsearch		CPU 0.22%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 346MB / 676kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.20%	Mem 195.4MiB / 256MiB = 76.34%	NetIO 672kB / 344MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.31%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 350MB / 682kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.27%	Mem 182.7MiB / 256MiB = 71.38%	NetIO 677kB / 346MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.02%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 356MB / 697kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.31%	Mem 185.8MiB / 256MiB = 72.57%	NetIO 698kB / 356MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.04%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 360MB / 704kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.59%	Mem 192.5MiB / 256MiB = 75.20%	NetIO 705kB / 360MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.01%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 367MB / 717kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.84%	Mem 189.5MiB / 256MiB = 74.03%	NetIO 718kB / 367MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.29%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 374MB / 732kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.42%	Mem 196.7MiB / 256MiB = 76.85%	NetIO 726kB / 370MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.03%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 377MB / 738kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.71%	Mem 204.2MiB / 256MiB = 79.78%	NetIO 739kB / 377MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.22%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 387MB / 759kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.53%	Mem 189.4MiB / 256MiB = 73.98%	NetIO 755kB / 385MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.33%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 391MB / 766kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.55%	Mem 193MiB / 256MiB = 75.38%	NetIO 761kB / 387MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.02%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 397MB / 780kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.84%	Mem 193.9MiB / 256MiB = 75.75%	NetIO 782kB / 397MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 401MB / 787kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.63%	Mem 201.9MiB / 256MiB = 78.87%	NetIO 788kB / 401MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 407MB / 800kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.59%	Mem 183.4MiB / 256MiB = 71.64%	NetIO 802kB / 407MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.26%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 415MB / 814kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.36%	Mem 186.3MiB / 256MiB = 72.76%	NetIO 808kB / 411MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 19.58MiB / 512MiB = 3.82%	NetIO 417MB / 820kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.87%	Mem 195.3MiB / 256MiB = 76.29%	NetIO 821kB / 417MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.28%	Mem 19.6MiB / 512MiB = 3.83%	NetIO 425MB / 833kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.45%	Mem 197.7MiB / 256MiB = 77.24%	NetIO 828kB / 421MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.03%	Mem 19.6MiB / 512MiB = 3.83%	NetIO 428MB / 838kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.07%	Mem 186MiB / 256MiB = 72.67%	NetIO 839kB / 428MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 19.6MiB / 512MiB = 3.83%	NetIO 435MB / 850kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.73%	Mem 194.6MiB / 256MiB = 76.01%	NetIO 852kB / 435MB	BlockIO 8.19kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 19.6MiB / 512MiB = 3.83%	NetIO 438MB / 857kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.85%	Mem 192.8MiB / 256MiB = 75.32%	NetIO 865kB / 442MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.38%	Mem 19.6MiB / 512MiB = 3.83%	NetIO 448MB / 877kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.93%	Mem 186.9MiB / 256MiB = 73.00%	NetIO 872kB / 446MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.01%	Mem 19.6MiB / 512MiB = 3.83%	NetIO 452MB / 884kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.94%	Mem 202.8MiB / 256MiB = 79.22%	NetIO 885kB / 452MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.03%	Mem 19.61MiB / 512MiB = 3.83%	NetIO 458MB / 896kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.75%	Mem 200.6MiB / 256MiB = 78.37%	NetIO 898kB / 458MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.03%	Mem 19.61MiB / 512MiB = 3.83%	NetIO 462MB / 904kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.58%	Mem 185.8MiB / 256MiB = 72.57%	NetIO 905kB / 462MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.06%	Mem 19.61MiB / 512MiB = 3.83%	NetIO 468MB / 917kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.83%	Mem 196.9MiB / 256MiB = 76.92%	NetIO 918kB / 468MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.28%	Mem 19.61MiB / 512MiB = 3.83%	NetIO 476MB / 931kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.66%	Mem 199.4MiB / 256MiB = 77.90%	NetIO 926kB / 472MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.01%	Mem 19.61MiB / 512MiB = 3.83%	NetIO 479MB / 937kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.06%	Mem 189MiB / 256MiB = 73.83%	NetIO 939kB / 479MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.28%	Mem 19.61MiB / 512MiB = 3.83%	NetIO 486MB / 950kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.50%	Mem 190.8MiB / 256MiB = 74.53%	NetIO 951kB / 486MB	BlockIO 12.3kB / 36.9kB
mockelasticsearch		CPU 0.34%	Mem 19.66MiB / 512MiB = 3.84%	NetIO 493MB / 963kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.23%	Mem 200.1MiB / 256MiB = 78.16%	NetIO 958kB / 489MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.02%	Mem 19.66MiB / 512MiB = 3.84%	NetIO 499MB / 976kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.49%	Mem 204.2MiB / 256MiB = 79.78%	NetIO 978kB / 499MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.01%	Mem 19.66MiB / 512MiB = 3.84%	NetIO 503MB / 984kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.98%	Mem 198.1MiB / 256MiB = 77.37%	NetIO 985kB / 503MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.26%	Mem 19.66MiB / 512MiB = 3.84%	NetIO 509MB / 995kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.89%	Mem 200.5MiB / 256MiB = 78.30%	NetIO 997kB / 509MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.01%	Mem 19.66MiB / 512MiB = 3.84%	NetIO 509MB / 995kB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 104.66%	Mem 208.8MiB / 256MiB = 81.55%	NetIO 997kB / 509MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.41%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 513MB / 1MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 103.28%	Mem 217MiB / 256MiB = 84.75%	NetIO 997kB / 509MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.27%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 519MB / 1.01MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.90%	Mem 218.9MiB / 256MiB = 85.51%	NetIO 1.01MB / 517MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.04%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 523MB / 1.02MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.51%	Mem 221.8MiB / 256MiB = 86.63%	NetIO 1.02MB / 523MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.05%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 529MB / 1.03MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.94%	Mem 208.3MiB / 256MiB = 81.39%	NetIO 1.04MB / 529MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.28%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 540MB / 1.05MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.55%	Mem 208.4MiB / 256MiB = 81.43%	NetIO 1.05MB / 537MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.02%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 544MB / 1.06MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.83%	Mem 203.8MiB / 256MiB = 79.62%	NetIO 1.06MB / 544MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.03%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 550MB / 1.07MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.56%	Mem 215.2MiB / 256MiB = 84.07%	NetIO 1.07MB / 550MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.24%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 560MB / 1.09MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.00%	Mem 218.4MiB / 256MiB = 85.30%	NetIO 1.09MB / 558MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.34%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 564MB / 1.1MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.05%	Mem 194.1MiB / 256MiB = 75.83%	NetIO 1.1MB / 560MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.21%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 570MB / 1.11MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.80%	Mem 194.9MiB / 256MiB = 76.12%	NetIO 1.11MB / 568MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.26%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 574MB / 1.12MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.55%	Mem 211MiB / 256MiB = 82.43%	NetIO 1.12MB / 574MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.21%	Mem 19.73MiB / 512MiB = 3.85%	NetIO 580MB / 1.13MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.48%	Mem 205.8MiB / 256MiB = 80.38%	NetIO 1.13MB / 578MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.01%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 584MB / 1.14MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.65%	Mem 201.2MiB / 256MiB = 78.61%	NetIO 1.14MB / 584MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.01%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 591MB / 1.15MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.90%	Mem 200.9MiB / 256MiB = 78.48%	NetIO 1.15MB / 591MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.27%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 598MB / 1.16MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.08%	Mem 206.6MiB / 256MiB = 80.70%	NetIO 1.16MB / 594MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.01%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 601MB / 1.17MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.74%	Mem 214.9MiB / 256MiB = 83.95%	NetIO 1.17MB / 601MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.34%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 609MB / 1.19MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.76%	Mem 197.1MiB / 256MiB = 77.01%	NetIO 1.18MB / 605MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.27%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 615MB / 1.2MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.70%	Mem 206.4MiB / 256MiB = 80.64%	NetIO 1.19MB / 611MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.21%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 621MB / 1.21MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.25%	Mem 211.9MiB / 256MiB = 82.77%	NetIO 1.21MB / 619MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.01%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 621MB / 1.21MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.39%	Mem 200.6MiB / 256MiB = 78.36%	NetIO 1.22MB / 625MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.02%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 631MB / 1.23MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.38%	Mem 203.9MiB / 256MiB = 79.66%	NetIO 1.23MB / 631MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.28%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 635MB / 1.24MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.51%	Mem 210.6MiB / 256MiB = 82.26%	NetIO 1.24MB / 635MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.06%	Mem 19.75MiB / 512MiB = 3.86%	NetIO 641MB / 1.25MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.74%	Mem 211.8MiB / 256MiB = 82.73%	NetIO 1.25MB / 641MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.01%	Mem 19.76MiB / 512MiB = 3.86%	NetIO 645MB / 1.26MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.74%	Mem 209.2MiB / 256MiB = 81.73%	NetIO 1.26MB / 645MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.03%	Mem 19.77MiB / 512MiB = 3.86%	NetIO 652MB / 1.27MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.65%	Mem 208.5MiB / 256MiB = 81.44%	NetIO 1.27MB / 652MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.32%	Mem 19.77MiB / 512MiB = 3.86%	NetIO 659MB / 1.28MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.24%	Mem 213.6MiB / 256MiB = 83.45%	NetIO 1.28MB / 656MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.03%	Mem 19.77MiB / 512MiB = 3.86%	NetIO 662MB / 1.29MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 102.38%	Mem 201.7MiB / 256MiB = 78.78%	NetIO 1.29MB / 662MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.29%	Mem 19.81MiB / 512MiB = 3.87%	NetIO 670MB / 1.31MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.59%	Mem 209.1MiB / 256MiB = 81.67%	NetIO 1.3MB / 666MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.28%	Mem 19.81MiB / 512MiB = 3.87%	NetIO 676MB / 1.32MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.35%	Mem 217.3MiB / 256MiB = 84.89%	NetIO 1.31MB / 672MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.23%	Mem 19.81MiB / 512MiB = 3.87%	NetIO 682MB / 1.33MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.23%	Mem 218.5MiB / 256MiB = 85.35%	NetIO 1.33MB / 680MB	BlockIO 16.4kB / 53.2kB
mockelasticsearch		CPU 0.24%	Mem 19.81MiB / 512MiB = 3.87%	NetIO 686MB / 1.34MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.76%	Mem 204.9MiB / 256MiB = 80.04%	NetIO 1.33MB / 682MB	BlockIO 16.4kB / 57.3kB
mockelasticsearch		CPU 0.01%	Mem 19.81MiB / 512MiB = 3.87%	NetIO 692MB / 1.35MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 101.90%	Mem 206.9MiB / 256MiB = 80.82%	NetIO 1.35MB / 692MB	BlockIO 16.4kB / 57.3kB
mockelasticsearch		CPU 0.34%	Mem 19.81MiB / 512MiB = 3.87%	NetIO 696MB / 1.36MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 100.56%	Mem 217.1MiB / 256MiB = 84.79%	NetIO 1.36MB / 696MB	BlockIO 16.4kB / 57.3kB
mockelasticsearch		CPU 0.31%	Mem 19.86MiB / 512MiB = 3.88%	NetIO 703MB / 1.37MB	BlockIO 65.5kB / 0B
fluentd-benchmark		CPU 35.08%	Mem 202.2MiB / 256MiB = 78.99%	NetIO 1.37MB / 703MB	BlockIO 16.4kB / 57.3kB
mockelasticsearch		CPU 0.03%	Mem 19.87MiB / 512MiB = 3.88%	NetIO 703MB / 1.37MB	BlockIO 65.5kB / 0B
```
</details>

### 500 files 1 MB each

```
STATS = overall bytes: 688216392, requests: 204, timestamps:
100 MB mark took 65714 milliseconds
200 MB mark took 68636 milliseconds
300 MB mark took 71989 milliseconds
400 MB mark took 71289 milliseconds
500 MB mark took 69174 milliseconds
600 MB mark took 72456 milliseconds
0 to 600 MB took 419261 milliseconds
```

<details>
  <summary>Docker Stats</summary>

```
mockelasticsearch		CPU 0.00%	Mem 1.582MiB / 512MiB = 0.31%	NetIO 650B / 0B	BlockIO 106kB / 0B
fluentd-benchmark		CPU 99.64%	Mem 50.5MiB / 256MiB = 19.73%	NetIO 1.03kB / 0B	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.00%	Mem 1.582MiB / 512MiB = 0.31%	NetIO 1.12kB / 0B	BlockIO 106kB / 0B
fluentd-benchmark		CPU 99.88%	Mem 112.4MiB / 256MiB = 43.91%	NetIO 1.19kB / 0B	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.17%	Mem 1.68MiB / 512MiB = 0.33%	NetIO 2.41kB / 1.78kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.80%	Mem 156.1MiB / 256MiB = 60.96%	NetIO 3.01kB / 1.22kB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.73%	Mem 6.445MiB / 512MiB = 1.26%	NetIO 3.9MB / 10kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.89%	Mem 161.9MiB / 256MiB = 63.26%	NetIO 3.01kB / 1.22kB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.04%	Mem 10.59MiB / 512MiB = 2.07%	NetIO 10.1MB / 23.1kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.90%	Mem 164.7MiB / 256MiB = 64.35%	NetIO 24.4kB / 10.1MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.03%	Mem 10.59MiB / 512MiB = 2.07%	NetIO 14MB / 29.5kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.04%	Mem 168.1MiB / 256MiB = 65.68%	NetIO 30.8kB / 14MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.01%	Mem 10.59MiB / 512MiB = 2.07%	NetIO 20.3MB / 42.5kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.99%	Mem 168.2MiB / 256MiB = 65.69%	NetIO 43.8kB / 20.3MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.38%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 28.1MB / 57.1kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.51%	Mem 162.5MiB / 256MiB = 63.49%	NetIO 51.1kB / 24.2MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.01%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 30.4MB / 63.1kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.51%	Mem 172MiB / 256MiB = 67.20%	NetIO 64.5kB / 30.4MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.29%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 38.2MB / 75.8kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.13%	Mem 178.6MiB / 256MiB = 69.77%	NetIO 70.8kB / 34.3MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.61%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 44.5MB / 89.7kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.84%	Mem 170.4MiB / 256MiB = 66.56%	NetIO 83.7kB / 40.6MB	BlockIO 8.19kB / 0B
mockelasticsearch		CPU 0.26%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 48.4MB / 95.4kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.13%	Mem 177.2MiB / 256MiB = 69.21%	NetIO 102kB / 50.7MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.24%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 54.6MB / 108kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.32%	Mem 181.2MiB / 256MiB = 70.79%	NetIO 110kB / 54.6MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.01%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 60.9MB / 122kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 103.10%	Mem 172.9MiB / 256MiB = 67.55%	NetIO 123kB / 60.9MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.02%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 64.8MB / 128kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.30%	Mem 173.9MiB / 256MiB = 67.93%	NetIO 129kB / 64.8MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.01%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 71MB / 140kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.16%	Mem 183.1MiB / 256MiB = 71.53%	NetIO 141kB / 71MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.01%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 74.9MB / 147kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.77%	Mem 187.4MiB / 256MiB = 73.20%	NetIO 155kB / 78.8MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.04%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 81.1MB / 159kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.45%	Mem 174.1MiB / 256MiB = 68.01%	NetIO 161kB / 81.1MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.04%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 88.9MB / 173kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.61%	Mem 183.4MiB / 256MiB = 71.66%	NetIO 174kB / 88.9MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.31%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 95.2MB / 187kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.94%	Mem 187.3MiB / 256MiB = 73.15%	NetIO 181kB / 91.3MB	BlockIO 8.19kB / 4.1kB
mockelasticsearch		CPU 0.21%	Mem 10.92MiB / 512MiB = 2.13%	NetIO 101MB / 201kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.03%	Mem 189.8MiB / 256MiB = 74.15%	NetIO 196kB / 99.1MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.01%	Mem 10.96MiB / 512MiB = 2.14%	NetIO 105MB / 208kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.44%	Mem 177MiB / 256MiB = 69.14%	NetIO 209kB / 105MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.02%	Mem 10.96MiB / 512MiB = 2.14%	NetIO 112MB / 220kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 104.48%	Mem 173.4MiB / 256MiB = 67.75%	NetIO 221kB / 112MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.35%	Mem 10.96MiB / 512MiB = 2.14%	NetIO 115MB / 227kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.13%	Mem 183.8MiB / 256MiB = 71.79%	NetIO 228kB / 115MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.03%	Mem 10.96MiB / 512MiB = 2.14%	NetIO 122MB / 241kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 104.45%	Mem 179.2MiB / 256MiB = 70.00%	NetIO 242kB / 122MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.27%	Mem 10.98MiB / 512MiB = 2.14%	NetIO 130MB / 254kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.30%	Mem 182.9MiB / 256MiB = 71.44%	NetIO 249kB / 126MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.02%	Mem 10.98MiB / 512MiB = 2.14%	NetIO 132MB / 260kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.08%	Mem 191.4MiB / 256MiB = 74.76%	NetIO 262kB / 132MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.01%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 140MB / 274kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.38%	Mem 201.1MiB / 256MiB = 78.55%	NetIO 276kB / 140MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.23%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 146MB / 288kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.25%	Mem 172.4MiB / 256MiB = 67.33%	NetIO 282kB / 142MB	BlockIO 8.19kB / 12.3kB
mockelasticsearch		CPU 0.18%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 152MB / 302kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.00%	Mem 174.9MiB / 256MiB = 68.32%	NetIO 297kB / 150MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.21%	Mem 15.16MiB / 512MiB = 2.96%	NetIO 156MB / 308kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.53%	Mem 185.7MiB / 256MiB = 72.56%	NetIO 310kB / 156MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.01%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 162MB / 321kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.50%	Mem 192.7MiB / 256MiB = 75.29%	NetIO 323kB / 162MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.02%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 166MB / 328kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.75%	Mem 173.6MiB / 256MiB = 67.79%	NetIO 330kB / 166MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.02%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 172MB / 341kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 104.14%	Mem 184.4MiB / 256MiB = 72.04%	NetIO 342kB / 172MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.25%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 180MB / 355kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.87%	Mem 183.8MiB / 256MiB = 71.79%	NetIO 349kB / 176MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.01%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 183MB / 362kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.69%	Mem 196.2MiB / 256MiB = 76.65%	NetIO 364kB / 183MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.43%	Mem 15.15MiB / 512MiB = 2.96%	NetIO 190MB / 376kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.63%	Mem 199.8MiB / 256MiB = 78.03%	NetIO 370kB / 186MB	BlockIO 8.19kB / 16.4kB
mockelasticsearch		CPU 0.03%	Mem 15.16MiB / 512MiB = 2.96%	NetIO 193MB / 383kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.03%	Mem 189.2MiB / 256MiB = 73.91%	NetIO 391kB / 197MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.28%	Mem 15.16MiB / 512MiB = 2.96%	NetIO 203MB / 402kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.03%	Mem 184.6MiB / 256MiB = 72.11%	NetIO 397kB / 201MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.34%	Mem 15.16MiB / 512MiB = 2.96%	NetIO 207MB / 409kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.13%	Mem 194.4MiB / 256MiB = 75.92%	NetIO 403kB / 203MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.30%	Mem 15.29MiB / 512MiB = 2.99%	NetIO 213MB / 422kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.55%	Mem 196.2MiB / 256MiB = 76.65%	NetIO 418kB / 211MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.01%	Mem 15.29MiB / 512MiB = 2.99%	NetIO 217MB / 428kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.68%	Mem 202.6MiB / 256MiB = 79.12%	NetIO 430kB / 217MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.21%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 223MB / 441kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 103.39%	Mem 192.4MiB / 256MiB = 75.14%	NetIO 443kB / 223MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.29%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 227MB / 448kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.84%	Mem 180.8MiB / 256MiB = 70.64%	NetIO 443kB / 223MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.23%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 233MB / 460kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.94%	Mem 186.6MiB / 256MiB = 72.89%	NetIO 456kB / 231MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.01%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 237MB / 465kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.07%	Mem 192.4MiB / 256MiB = 75.16%	NetIO 467kB / 237MB	BlockIO 8.19kB / 20.5kB
mockelasticsearch		CPU 0.02%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 243MB / 479kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 105.84%	Mem 178MiB / 256MiB = 69.55%	NetIO 480kB / 243MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.03%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 247MB / 486kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.02%	Mem 183.1MiB / 256MiB = 71.54%	NetIO 488kB / 247MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.02%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 254MB / 498kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 104.32%	Mem 191.6MiB / 256MiB = 74.85%	NetIO 500kB / 254MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.05%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 257MB / 506kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.71%	Mem 197.4MiB / 256MiB = 77.13%	NetIO 514kB / 261MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.31%	Mem 15.3MiB / 512MiB = 2.99%	NetIO 268MB / 526kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.86%	Mem 203.5MiB / 256MiB = 79.51%	NetIO 521kB / 264MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.19%	Mem 15.78MiB / 512MiB = 3.08%	NetIO 274MB / 536kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.73%	Mem 190.9MiB / 256MiB = 74.59%	NetIO 533kB / 271MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.56%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 278MB / 545kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.17%	Mem 191.5MiB / 256MiB = 74.81%	NetIO 539kB / 274MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.29%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 284MB / 558kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.80%	Mem 193.9MiB / 256MiB = 75.74%	NetIO 553kB / 282MB	BlockIO 8.19kB / 28.7kB
mockelasticsearch		CPU 0.35%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 288MB / 566kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.41%	Mem 201.9MiB / 256MiB = 78.87%	NetIO 568kB / 288MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.02%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 294MB / 579kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 103.48%	Mem 208.2MiB / 256MiB = 81.34%	NetIO 581kB / 294MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 298MB / 586kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.29%	Mem 191.1MiB / 256MiB = 74.65%	NetIO 588kB / 298MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.04%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 304MB / 598kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 105.69%	Mem 189MiB / 256MiB = 73.84%	NetIO 599kB / 304MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 308MB / 604kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.21%	Mem 194.1MiB / 256MiB = 75.84%	NetIO 606kB / 308MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.03%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 314MB / 616kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 107.00%	Mem 204.9MiB / 256MiB = 80.05%	NetIO 618kB / 314MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.04%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 318MB / 624kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.46%	Mem 194.8MiB / 256MiB = 76.08%	NetIO 632kB / 322MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.01%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 325MB / 637kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.48%	Mem 195.4MiB / 256MiB = 76.34%	NetIO 639kB / 325MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.28%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 332MB / 652kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.34%	Mem 201MiB / 256MiB = 78.50%	NetIO 646kB / 328MB	BlockIO 12.3kB / 32.8kB
mockelasticsearch		CPU 0.35%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 339MB / 665kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.96%	Mem 206.2MiB / 256MiB = 80.56%	NetIO 660kB / 335MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.37%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 342MB / 673kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.17%	Mem 212.9MiB / 256MiB = 83.17%	NetIO 680kB / 345MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.28%	Mem 15.31MiB / 512MiB = 2.99%	NetIO 349MB / 686kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.81%	Mem 188.9MiB / 256MiB = 73.78%	NetIO 680kB / 345MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.24%	Mem 15.68MiB / 512MiB = 3.06%	NetIO 355MB / 698kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.05%	Mem 197.8MiB / 256MiB = 77.27%	NetIO 695kB / 353MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.31%	Mem 15.32MiB / 512MiB = 2.99%	NetIO 359MB / 705kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.23%	Mem 200.2MiB / 256MiB = 78.20%	NetIO 700kB / 355MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.20%	Mem 15.33MiB / 512MiB = 2.99%	NetIO 365MB / 718kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.61%	Mem 203.1MiB / 256MiB = 79.35%	NetIO 714kB / 363MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.26%	Mem 15.32MiB / 512MiB = 2.99%	NetIO 369MB / 725kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.93%	Mem 192.6MiB / 256MiB = 75.23%	NetIO 720kB / 365MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.01%	Mem 15.32MiB / 512MiB = 2.99%	NetIO 373MB / 732kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 106.91%	Mem 201.6MiB / 256MiB = 78.74%	NetIO 740kB / 375MB	BlockIO 12.3kB / 41kB
mockelasticsearch		CPU 0.40%	Mem 15.33MiB / 512MiB = 2.99%	NetIO 379MB / 745kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.66%	Mem 207.1MiB / 256MiB = 80.89%	NetIO 746kB / 379MB	BlockIO 12.3kB / 45.1kB
mockelasticsearch		CPU 0.03%	Mem 15.33MiB / 512MiB = 2.99%	NetIO 385MB / 757kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 104.69%	Mem 196MiB / 256MiB = 76.54%	NetIO 759kB / 385MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.01%	Mem 15.34MiB / 512MiB = 3.00%	NetIO 389MB / 764kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.81%	Mem 206.5MiB / 256MiB = 80.67%	NetIO 765kB / 389MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.01%	Mem 11.65MiB / 512MiB = 2.28%	NetIO 396MB / 777kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 106.17%	Mem 210.8MiB / 256MiB = 82.36%	NetIO 779kB / 396MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.29%	Mem 11.65MiB / 512MiB = 2.28%	NetIO 403MB / 791kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.27%	Mem 206.7MiB / 256MiB = 80.73%	NetIO 785kB / 399MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.02%	Mem 11.66MiB / 512MiB = 2.28%	NetIO 406MB / 797kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.56%	Mem 201.6MiB / 256MiB = 78.75%	NetIO 798kB / 406MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.33%	Mem 11.66MiB / 512MiB = 2.28%	NetIO 413MB / 811kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.06%	Mem 203.8MiB / 256MiB = 79.61%	NetIO 806kB / 410MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.01%	Mem 11.66MiB / 512MiB = 2.28%	NetIO 416MB / 818kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.93%	Mem 213MiB / 256MiB = 83.19%	NetIO 819kB / 416MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.01%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 424MB / 831kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.06%	Mem 203.4MiB / 256MiB = 79.44%	NetIO 833kB / 424MB	BlockIO 12.3kB / 49.2kB
mockelasticsearch		CPU 0.29%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 430MB / 844kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.34%	Mem 204.3MiB / 256MiB = 79.82%	NetIO 838kB / 426MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.22%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 436MB / 856kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.22%	Mem 205.8MiB / 256MiB = 80.41%	NetIO 852kB / 434MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.09%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 440MB / 863kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.47%	Mem 202.5MiB / 256MiB = 79.10%	NetIO 863kB / 439MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.02%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 446MB / 876kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.97%	Mem 202.2MiB / 256MiB = 78.98%	NetIO 877kB / 446MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.02%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 450MB / 882kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.69%	Mem 208.9MiB / 256MiB = 81.61%	NetIO 884kB / 450MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.01%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 456MB / 895kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 103.78%	Mem 214.5MiB / 256MiB = 83.80%	NetIO 897kB / 456MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.02%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 460MB / 901kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.36%	Mem 202.1MiB / 256MiB = 78.95%	NetIO 903kB / 460MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.05%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 467MB / 914kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 106.86%	Mem 205.8MiB / 256MiB = 80.40%	NetIO 916kB / 467MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.30%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 474MB / 929kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.86%	Mem 214MiB / 256MiB = 83.61%	NetIO 923kB / 470MB	BlockIO 12.3kB / 57.3kB
mockelasticsearch		CPU 0.03%	Mem 11.71MiB / 512MiB = 2.29%	NetIO 477MB / 935kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 103.17%	Mem 198.5MiB / 256MiB = 77.52%	NetIO 937kB / 477MB	BlockIO 12.3kB / 61.4kB
mockelasticsearch		CPU 0.01%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 484MB / 949kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.07%	Mem 207MiB / 256MiB = 80.86%	NetIO 950kB / 484MB	BlockIO 12.3kB / 61.4kB
mockelasticsearch		CPU 0.29%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 491MB / 961kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.04%	Mem 211.1MiB / 256MiB = 82.48%	NetIO 957kB / 487MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.18%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 497MB / 975kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.14%	Mem 213.1MiB / 256MiB = 83.26%	NetIO 970kB / 495MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.36%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 501MB / 982kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.55%	Mem 203.8MiB / 256MiB = 79.60%	NetIO 976kB / 497MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.19%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 507MB / 994kB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.07%	Mem 204.5MiB / 256MiB = 79.86%	NetIO 990kB / 505MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.04%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 511MB / 1MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.77%	Mem 217.4MiB / 256MiB = 84.93%	NetIO 1MB / 511MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.29%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 517MB / 1.01MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 107.02%	Mem 209MiB / 256MiB = 81.65%	NetIO 1.02MB / 517MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.01%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 521MB / 1.02MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.99%	Mem 210.7MiB / 256MiB = 82.32%	NetIO 1.02MB / 521MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.02%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 527MB / 1.03MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 109.20%	Mem 218.1MiB / 256MiB = 85.19%	NetIO 1.04MB / 527MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.01%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 531MB / 1.04MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.30%	Mem 205.8MiB / 256MiB = 80.38%	NetIO 1.04MB / 531MB	BlockIO 12.3kB / 65.5kB
mockelasticsearch		CPU 0.02%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 538MB / 1.05MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 108.07%	Mem 210.7MiB / 256MiB = 82.30%	NetIO 1.05MB / 538MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.01%	Mem 11.7MiB / 512MiB = 2.29%	NetIO 541MB / 1.06MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.52%	Mem 214.6MiB / 256MiB = 83.85%	NetIO 1.06MB / 541MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.24%	Mem 11.71MiB / 512MiB = 2.29%	NetIO 548MB / 1.07MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.76%	Mem 205.1MiB / 256MiB = 80.11%	NetIO 1.07MB / 548MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.01%	Mem 11.71MiB / 512MiB = 2.29%	NetIO 552MB / 1.08MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.63%	Mem 210.5MiB / 256MiB = 82.22%	NetIO 1.08MB / 552MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.01%	Mem 11.71MiB / 512MiB = 2.29%	NetIO 558MB / 1.09MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 107.03%	Mem 205MiB / 256MiB = 80.09%	NetIO 1.09MB / 558MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.01%	Mem 11.71MiB / 512MiB = 2.29%	NetIO 562MB / 1.1MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.15%	Mem 207.7MiB / 256MiB = 81.13%	NetIO 1.11MB / 566MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.01%	Mem 11.78MiB / 512MiB = 2.30%	NetIO 568MB / 1.11MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.86%	Mem 216.9MiB / 256MiB = 84.74%	NetIO 1.11MB / 568MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.02%	Mem 11.78MiB / 512MiB = 2.30%	NetIO 576MB / 1.12MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.58%	Mem 226.3MiB / 256MiB = 88.39%	NetIO 1.13MB / 576MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.18%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 580MB / 1.13MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.75%	Mem 215.5MiB / 256MiB = 84.19%	NetIO 1.13MB / 578MB	BlockIO 12.3kB / 73.7kB
mockelasticsearch		CPU 0.04%	Mem 11.78MiB / 512MiB = 2.30%	NetIO 586MB / 1.14MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.12%	Mem 216MiB / 256MiB = 84.37%	NetIO 1.14MB / 586MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.31%	Mem 11.78MiB / 512MiB = 2.30%	NetIO 592MB / 1.16MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.12%	Mem 224.5MiB / 256MiB = 87.69%	NetIO 1.15MB / 588MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.20%	Mem 12.43MiB / 512MiB = 2.43%	NetIO 598MB / 1.17MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.67%	Mem 210.3MiB / 256MiB = 82.14%	NetIO 1.16MB / 596MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.32%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 602MB / 1.17MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.53%	Mem 218.6MiB / 256MiB = 85.40%	NetIO 1.17MB / 598MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.16%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 609MB / 1.19MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.89%	Mem 219.5MiB / 256MiB = 85.76%	NetIO 1.18MB / 606MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.05%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 612MB / 1.19MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.92%	Mem 216.5MiB / 256MiB = 84.56%	NetIO 1.19MB / 612MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.19%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 619MB / 1.2MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.41%	Mem 217.3MiB / 256MiB = 84.87%	NetIO 1.2MB / 616MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.02%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 623MB / 1.21MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.70%	Mem 224.6MiB / 256MiB = 87.73%	NetIO 1.21MB / 623MB	BlockIO 16.4kB / 81.9kB
mockelasticsearch		CPU 0.26%	Mem 11.79MiB / 512MiB = 2.30%	NetIO 629MB / 1.23MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 105.79%	Mem 216MiB / 256MiB = 84.38%	NetIO 1.23MB / 629MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.01%	Mem 11.82MiB / 512MiB = 2.31%	NetIO 633MB / 1.23MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.37%	Mem 226.6MiB / 256MiB = 88.51%	NetIO 1.23MB / 633MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.03%	Mem 11.82MiB / 512MiB = 2.31%	NetIO 637MB / 1.24MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 107.97%	Mem 218.9MiB / 256MiB = 85.50%	NetIO 1.25MB / 639MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.04%	Mem 11.82MiB / 512MiB = 2.31%	NetIO 643MB / 1.25MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.04%	Mem 220.3MiB / 256MiB = 86.07%	NetIO 1.25MB / 643MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.01%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 649MB / 1.27MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 105.31%	Mem 226.1MiB / 256MiB = 88.32%	NetIO 1.27MB / 649MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.01%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 653MB / 1.27MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.72%	Mem 215.1MiB / 256MiB = 84.03%	NetIO 1.27MB / 653MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.03%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 659MB / 1.29MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 108.91%	Mem 220.5MiB / 256MiB = 86.15%	NetIO 1.29MB / 659MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.01%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 663MB / 1.29MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 102.15%	Mem 230.4MiB / 256MiB = 89.99%	NetIO 1.3MB / 667MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.01%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 669MB / 1.3MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 107.24%	Mem 221.7MiB / 256MiB = 86.61%	NetIO 1.31MB / 669MB	BlockIO 16.4kB / 90.1kB
mockelasticsearch		CPU 0.34%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 677MB / 1.32MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.62%	Mem 229.9MiB / 256MiB = 89.81%	NetIO 1.31MB / 673MB	BlockIO 20.5kB / 102kB
mockelasticsearch		CPU 0.01%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 680MB / 1.32MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 101.92%	Mem 234.7MiB / 256MiB = 91.69%	NetIO 1.33MB / 680MB	BlockIO 20.5kB / 102kB
mockelasticsearch		CPU 0.26%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 687MB / 1.34MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 100.75%	Mem 222.8MiB / 256MiB = 87.01%	NetIO 1.34MB / 687MB	BlockIO 20.5kB / 102kB
mockelasticsearch		CPU 0.08%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 690MB / 1.34MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 4.38%	Mem 226.4MiB / 256MiB = 88.42%	NetIO 1.34MB / 690MB	BlockIO 20.5kB / 102kB
mockelasticsearch		CPU 0.04%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 690MB / 1.34MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 5.23%	Mem 227.1MiB / 256MiB = 88.71%	NetIO 1.34MB / 690MB	BlockIO 20.5kB / 102kB
mockelasticsearch		CPU 0.08%	Mem 15.63MiB / 512MiB = 3.05%	NetIO 690MB / 1.34MB	BlockIO 106kB / 0B
fluentd-benchmark		CPU 7.34%	Mem 227.7MiB / 256MiB = 88.95%	NetIO 1.34MB / 690MB	BlockIO 20.5kB / 102kB
fluentd-benchmark		CPU 3.92%	Mem 228.4MiB / 256MiB = 89.24%	NetIO 1.34MB / 690MB	BlockIO 20.5kB / 106kB
```
</details>


## Bench 2 - Version 1.14-1

* Docker container CPU limit: 0.5
* Docker container memory limit: 512 MB

### 1 file 500 MB

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 697759611, requests: 205, timestamps:
100 MB mark took 142584 milliseconds
200 MB mark took 146933 milliseconds
300 MB mark took 144257 milliseconds
400 MB mark took 145342 milliseconds
500 MB mark took 150089 milliseconds
600 MB mark took 149173 milliseconds
0 to 600 MB took 878381 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 689759605, requests: 202, timestamps:
100 MB mark took 21905 milliseconds
200 MB mark took 21705 milliseconds
300 MB mark took 21186 milliseconds
400 MB mark took 21518 milliseconds
500 MB mark took 21604 milliseconds
600 MB mark took 22568 milliseconds
0 to 600 MB took 130488 milliseconds
```
</details>

### 500 files 1 MB each

<details>
  <summary>Macbook Pro M1 Stats</summary>

```
STATS = overall bytes: 697987803, requests: 207, timestamps:
100 MB mark took 142378 milliseconds
200 MB mark took 150197 milliseconds
300 MB mark took 154336 milliseconds
400 MB mark took 151696 milliseconds
500 MB mark took 148872 milliseconds
600 MB mark took 152913 milliseconds
0 to 600 MB took 900395 milliseconds
```
</details>

<details>
  <summary>Linux Intel Xeon Stats</summary>

```
STATS = overall bytes: 689983679, requests: 207, timestamps:
100 MB mark took 22670 milliseconds
200 MB mark took 21615 milliseconds
300 MB mark took 22903 milliseconds
400 MB mark took 21896 milliseconds
500 MB mark took 23578 milliseconds
600 MB mark took 22534 milliseconds
0 to 600 MB took 135199 milliseconds
```
</details>
