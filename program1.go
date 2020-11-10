package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Oops"))
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/body", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello Body")
	})
	http.ListenAndServe(":9090", nil)
}
