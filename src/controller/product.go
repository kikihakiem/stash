package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kikihakiem/stash/go/simple-crud/repository"
)

func ProductController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product := repository.GetProductBySKU(vars["sku"])
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}
