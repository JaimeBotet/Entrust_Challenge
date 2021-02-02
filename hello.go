package main

import (
	"net/http"
	"fmt"
	"time"
)

func f(from string,i int) {
	fmt.Println(from, " Sleeping...", i)
	time.Sleep(time.Second)
}


func main() {
	// create a new 'ServeMux'
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
}