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
	r.HandleFunc("/{sku}", c.GetBySKU)
}

func (c *ProductController) GetBySKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product := repository.GetProductBySKU(vars["sku"])
	json.NewEncoder(w).Encode(product)
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
	}

	product.ID = 3
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}
