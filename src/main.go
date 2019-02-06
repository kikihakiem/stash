package main

import (
	"github.com/gorilla/mux"
	"github.com/kikihakiem/stash/go/simple-crud/controller"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("/{sku}", controller.ProductController)
}
