package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kikihakiem/stash/go/simple-crud/repository"
)

type ProductController struct {
	db   *sql.DB
	logf func(string, ...interface{})
}

// NewProductController accepts *sql.DB, so we can replace it with stub in testing
func NewProductController(db *sql.DB, logf func(string, ...interface{})) *ProductController {
	return &ProductController{db: db, logf: logf}
}

// Handle handles subroutes/path
func (c *ProductController) Handle(r *mux.Router) {
	r.HandleFunc("", c.Create).Methods("POST")
	r.HandleFunc("", c.GetAll).Methods("GET")
	r.HandleFunc("/{sku}", c.GetBySKU).Methods("GET")
	r.HandleFunc("/{sku}", c.Update).Methods("PATCH")
	r.HandleFunc("/{sku}", c.Delete).Methods("DELETE")
}

func (c *ProductController) GetBySKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["sku"] == "" {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}
	product := repository.GetProductBySKU(vars["sku"])
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	products := repository.GetAll()
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		c.logf("error encoding response body: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var product repository.Product
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&product)
	if err != nil {
		c.logf("error decoding request body: %v", err)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	product.ID = 3
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["sku"] == "" {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	var product repository.Product
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&product)
	if err != nil {
		c.logf("error decoding request body: %v", err)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	product.Category = "Men - Casual - Pants"
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["sku"] == "" {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}
	w.WriteHeader(204)
}
