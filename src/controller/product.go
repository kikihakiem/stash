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

	product, err := repository.GetProductBySKU(c.db, vars["sku"])
	if err != nil {
		c.logf("error getting product by SKU: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetAllProduct(c.db)
	if err != nil {
		c.logf("error getting all products: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(products)
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

	product, err = repository.CreateProduct(c.db, product)
	if err != nil {
		c.logf("error creating product: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

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

	product, err = repository.UpdateProduct(c.db, product)
	if err != nil {
		c.logf("error updating product: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["sku"] == "" {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	err := repository.DeleteProduct(c.db, vars["sku"])
	if err != nil {
		c.logf("error deleting product: %v", err)
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
}
