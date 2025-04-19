package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "test")
	})
	http.ListenAndServe(":8080", r)
}
