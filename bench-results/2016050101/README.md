

- Google Cloud Compute Engine
- 8 vCPUs + 16GB Memory + 50GB SSD
- 1 machine(client) of 16 vCPUs + 30GB Memory + 50GB SSD
- Ubuntu 15.10
- Go master branch on 2016-04-30
- etcd v3



<br><br><hr>
##### Write 2K keys, 1-client, 1-conn, 8-byte key with random texts(<300KB) (etcd v3)

<img src="https://storage.googleapis.com/dbtester-results/2016050101/01-avg-latency-ms.png" alt="2016050101/01-avg-latency-ms">

<img src="https://storage.googleapis.com/dbtester-results/2016050101/01-throughput.png" alt="2016050101/01-throughput">

<img src="https://storage.googleapis.com/dbtester-results/2016050101/01-avg-cpu.png" alt="2016050101/01-avg-cpu">

<img src="https://storage.googleapis.com/dbtester-results/2016050101/01-avg-memory.png" alt="2016050101/01-avg-memory">



<br><br><hr>
##### Write 3K keys, 1-client, 1-conn, 8-byte key with random texts(<300KB) (etcd v3)

<img src="https://storage.googleapis.com/dbtester-results/2016050101/02-avg-latency-ms.png" alt="2016050101/02-avg-latency-ms">

<img src="https://storage.googleapis.com/dbtester-results/2016050101/02-throughput.png" alt="2016050101/02-throughput">

<img src="https://storage.googleapis.com/dbtester-results/2016050101/02-avg-cpu.png" alt="2016050101/02-avg-cpu">

<img src="https://storage.googleapis.com/dbtester-results/2016050101/02-avg-memory.png" alt="2016050101/02-avg-memory">



