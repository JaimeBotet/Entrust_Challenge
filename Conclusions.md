| Mode             | Concurrency | Complete Requests | Failed Requests | Requests per second | Time per Request(mean) | TPR (mean, across all concurrent requests) |
| :--------------- | :---------: | :---------------: | :-------------: | :-----------------: | :--------------------: | -----------------------------------------: |
| -n 1000          |      1      |       1000        |        0        |       1817.17       |         0.550          |                                      0.550 |
| -n 1000 -c 10    |     10      |       1000        |        0        |       2104.59       |         4.752          |                                      0.475 |
| -n 1000 -c 10 -k |     10      |       1000        |        0        |      11797.04       |         0.848          |                                      0.085 |

```
ab -n 1000 http://localhost:9000/go/async

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       2
Processing:     0    0   0.4      0       2
Waiting:        0    0   0.3      0       1
Total:          0    0   0.5      0       2
```

#####

```
ab -n 1000 -c 10 http://localhost:9000/go/async

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       2
Processing:     1    4   1.3      4       9
Waiting:        0    3   1.4      3       8
Total:          1    5   1.3      4      10
```

#####

```
ab -n 1000 -c 10 -k http://localhost:9000/go/async

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     0    1   1.7      0      13
Waiting:        0    1   1.7      0      13
Total:          0    1   1.7      0      13
```
