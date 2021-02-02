# Go AB

Testing Go Server with Apache Workbench.

Using GoRoutines to see concurrency.

## Basic Go Server with Different Routes

Here is the basic implementation of a Go Server with a router of 2 different paths:

    mux := http.NewServeMux()
    //Handle '/sync' route
    mux.HandleFunc("/sync", func(res http.ResponseWriter, req *http.Request) {
    	for i := 0; i < 3; i++ {
    		f("sync", i)
    	}
    	fmt.Fprint(res, "Finished sync request")
    })

    //Handle '/go/async' route
    mux.HandleFunc("/go/async", func(res http.ResponseWriter, req *http.Request) {
    	for i := 0; i < 3; i++ {
    		go f("async",i)
    	}
    	fmt.Fprint(res, "Finished async request")
    })

    // listen and serve using 'ServeMux'
    http.ListenAndServe(":9000", mux)

## AB Runs

Here I executed different combinations of **ab** on the server I created to see the performance depending on the _concurrency_.

| Mode             | Concurrency | Complete Requests | Failed Requests | Requests per second | Time per Request(mean) | TPR (mean, across all concurrent requests) |
| ---------------- | ----------- | ----------------- | --------------- | ------------------- | ---------------------- | ------------------------------------------ |
| -n 1000          | 1           | 1000              | 0               | 1817.17             | 0.550                  | 0.550                                      |
| -n 1000 -c 10    | 10          | 1000              | 0               | 2104.59             | 4.752                  | 0.475                                      |
| -n 1000 -c 50    | 50          | 1000              | 0               | 2348.18             | 21.293                 | 0.426                                      |
| -n 1000 -c 100   | 100         | 1000              | 0               | 2258.29             | 44.281                 | 0.443                                      |
| -n 1000 -c 10 -k | 10          | 1000              | 0               | 11797.04            | 0.848                  | 0.085                                      |

From this Data we can inferr following conclusions:

- Referring only to _concurrency_ I can only conclude that increasing indefinitely the concurrency level wont outcome an increase of the number of requests per second, this will come limited with the own processors of the PC. In my case (I used windows) I modified the configuration so my PC would make use of the 8 cores.

- As we can see from the above data, with concurrency 50 we obtain better results in terms of _Requests Per Second_ than the ones with concurrency of 100.

- It also makes sense that as long as our processors can handle it, the concurrency will increase the number of _Requests Per Second_.

- It is important also to take into account the difference when using the parameter _-k_. This _KeepAlive_ feature saves us time when opening/closing the channels, keeping them 'Alive' so there is no time wasted in the closing and creation of channels for the secuential processes.
