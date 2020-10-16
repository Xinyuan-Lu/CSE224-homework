package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mut sync.Mutex

func main() {
	http.HandleFunc("/", handler)               // each request calls handler
	http.HandleFunc("/gendata", handlerGendata) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	mut.Lock()
	fmt.Fprint(w, "Hello, There.")
	mut.Unlock()
}
func handlerGendata(w http.ResponseWriter, r *http.Request) {
	mut.Lock()
	log.Printf("inbound request: %q, header: %q", r.RequestURI, r.Header)
	var y string
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "numBytes" {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprint(w, "Error in Converting the input number")
			}
			y = strings.Repeat(".", i)
			fmt.Fprintf(w, "%s", y)
			return
		}
	}
	mut.Unlock()
}
