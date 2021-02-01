ab -n 1000 http://localhost/PHP_FileSystem/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software: Apache/2.4.43
Server Hostname: localhost
Server Port: 80

Document Path: /PHP_FileSystem/
Document Length: 8536 bytes

Concurrency Level: 1
Time taken for tests: 1.463 seconds
Complete requests: 1000
Failed requests: 0
Total transferred: 8732000 bytes
HTML transferred: 8536000 bytes
Requests per second: 683.42 [#/sec] (mean)
Time per request: 1.463 [ms] (mean)
Time per request: 1.463 [ms] (mean, across all concurrent requests)
Transfer rate: 5827.76 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max
Connect: 0 0 0.4 0 1
Processing: 0 1 0.4 1 5
Waiting: 0 1 0.4 1 3
Total: 0 1 0.6 1 6

Percentage of the requests served within a certain time (ms)
50% 1
66% 1
75% 2
80% 2
90% 2
95% 2
98% 3
99% 3
100% 6 (longest request)

#####

PS D:\xampp\htdocs\Entrust_challenge> ab -n 1000 http://localhost:9000/hello/golang
This is ApacheBench, Version 2.3 <$Revision: 1874286 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software:
Server Hostname: localhost
Server Port: 9000

Document Path: /hello/golang
Document Length: 13 bytes

Concurrency Level: 1
Time taken for tests: 0.496 seconds
Complete requests: 1000
Failed requests: 0
Total transferred: 130000 bytes
HTML transferred: 13000 bytes
Requests per second: 2017.48 [#/sec] (mean)
Time per request: 0.496 [ms] (mean)
Time per request: 0.496 [ms] (mean, across all concurrent requests)
Transfer rate: 256.13 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max
Connect: 0 0 0.4 0 1
Processing: 0 0 0.4 0 2
Waiting: 0 0 0.3 0 1
Total: 0 0 0.5 0 2

Percentage of the requests served within a certain time (ms)
50% 0
66% 1
75% 1
80% 1
90% 1
95% 1
98% 1
99% 1
100% 2 (longest request)

#####

PS D:\xampp\htdocs\Entrust_challenge> ab -n 1000 -c 10 http://localhost/PHP_FileSystem/
This is ApacheBench, Version 2.3 <$Revision: 1874286 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software: Apache/2.4.43
Server Hostname: localhost
Server Port: 80

Document Path: /PHP_FileSystem/
Document Length: 8536 bytes

Concurrency Level: 10
Time taken for tests: 0.475 seconds
Complete requests: 1000
Failed requests: 0
Total transferred: 8732000 bytes
HTML transferred: 8536000 bytes
Requests per second: 2106.47 [#/sec] (mean)
Time per request: 4.747 [ms] (mean)
Time per request: 0.475 [ms] (mean, across all concurrent requests)
Transfer rate: 17962.59 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max
Connect: 0 0 0.4 0 2
Waiting: 1 3 1.3 3 14
Total: 1 5 1.4 4 17

Percentage of the requests served within a certain time (ms)
50% 4
66% 5
75% 5
80% 5
90% 7
95% 7
98% 8
99% 8
100% 17 (longest request)

#####

PS D:\xampp\htdocs\Entrust_challenge> ab -n 1000 -c 10 http://localhost:9000/hello/golang
This is ApacheBench, Version 2.3 <$Revision: 1874286 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software:
Server Hostname: localhost
Server Port: 9000

Document Path: /hello/golang
Document Length: 13 bytes

Concurrency Level: 10
Time taken for tests: 0.422 seconds
Complete requests: 1000
Failed requests: 0
Total transferred: 130000 bytes
HTML transferred: 13000 bytes
Requests per second: 2370.47 [#/sec] (mean)
Time per request: 4.219 [ms] (mean)
Time per request: 0.422 [ms] (mean, across all concurrent requests)
Transfer rate: 300.94 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max
Connect: 0 0 0.4 0 2
Processing: 0 4 1.5 3 12
Waiting: 0 3 1.5 2 12
Total: 0 4 1.5 4 13

Percentage of the requests served within a certain time (ms)
50% 4
66% 4
75% 4
80% 5
90% 6
95% 7
98% 8
99% 9
100% 13 (longest request)

#####

#####

PS D:\xampp\htdocs\Entrust_challenge> ab -n 1000 -c 10 -k http://localhost:9000/hello/golang
This is ApacheBench, Version 2.3 <$Revision: 1874286 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software:
Server Hostname: localhost
Server Port: 9000

Document Path: /hello/golang
Document Length: 13 bytes

Concurrency Level: 10
Time taken for tests: 0.088 seconds
Complete requests: 1000
Failed requests: 0
Keep-Alive requests: 1000
Total transferred: 154000 bytes
HTML transferred: 13000 bytes
Requests per second: 11396.27 [#/sec] (mean)
Time per request: 0.877 [ms] (mean)
Time per request: 0.088 [ms] (mean, across all concurrent requests)
Transfer rate: 1713.89 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max
Connect: 0 0 0.1 0 1
Processing: 0 1 1.4 0 12
Waiting: 0 1 1.4 0 12
Total: 0 1 1.4 0 12

Percentage of the requests served within a certain time (ms)
50% 0
66% 1
75% 1
80% 1
90% 1
95% 3
98% 6
99% 8
100% 12 (longest request)

#####

PS D:\xampp\htdocs\Entrust_challenge> ab -n 1000 -c 10 -k http://localhost/PHP_FileSystem/  
This is ApacheBench, Version 2.3 <$Revision: 1874286 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software: Apache/2.4.43
Server Hostname: localhost
Server Port: 80

Document Path: /PHP_FileSystem/
Document Length: 8536 bytes

Concurrency Level: 10
Time taken for tests: 0.475 seconds
Complete requests: 1000
Failed requests: 0
Keep-Alive requests: 0
Total transferred: 8732000 bytes
HTML transferred: 8536000 bytes
Requests per second: 2106.45 [#/sec] (mean)
Time per request: 4.747 [ms] (mean)
Time per request: 0.475 [ms] (mean, across all concurrent requests)
Transfer rate: 17962.44 [Kbytes/sec] received

Connection Times (ms)
min mean[+/-sd] median max
Connect: 0 0 0.4 0 2
Processing: 1 4 1.3 4 12
Waiting: 1 3 1.2 3 10
Total: 1 5 1.3 4 12

Percentage of the requests served within a certain time (ms)
50% 4
66% 5
75% 5
80% 5
90% 6
95% 7
98% 8
99% 10
100% 12 (longest request)

#####
