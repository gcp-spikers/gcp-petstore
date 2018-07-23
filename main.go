package main

import (
	"io"
	"net/http"
	"flag"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "port number")
	flag.Parse()

	fmt.Printf("Starting server on port %d...", port)
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); 
	
	if err != nil {
		panic(err)
	}
}
