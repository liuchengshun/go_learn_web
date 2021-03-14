package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path:%q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for i, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q \n", i, v)
	}
}