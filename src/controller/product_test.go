package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gorilla/mux"

	"github.com/kikihakiem/stash/go/simple-crud/common"
	"github.com/kikihakiem/stash/go/simple-crud/repository"
)

var (
	db     *sql.DB
	initDB sync.Once
)

func getDB() *sql.DB {
	initDB.Do(func() {
		if db == nil {
			var err error
			db, err = common.GetDB("127.0.0.1", "3306",
				"dbuser", "y3T aN0tH3r 5tr0nG P4s5WoRd", "product_db")
			if err != nil {
				panic(err)
			}
		}
	})

	return db
}

func TestGETSingleProduct(t *testing.T) {
	t.Run("returns product details", func(t *testing.T) {
		sku := "abc123"
		// setup product
		p := repository.Product{
			SKU:      sku,
			Name:     "Cardinal Jeans",
			Category: "Men - Casual - Trousers",
		}
		repository.CreateProduct(getDB(), p)

		// setup HTTP request
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", sku), nil)
		request = mux.SetURLVars(request, map[string]string{"sku": sku})
		response := httptest.NewRecorder()

		productController := NewProductController(getDB(), log.Printf)
		productController.GetBySKU(response, request)

		respCode := response.Code
		expectedRespCode := 200
		assertEqual(t, respCode, expectedRespCode)

		var product repository.Product
		err := json.NewDecoder(response.Body).Decode(&product)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into Product, '%v'", response.Body, err)
		}

		// got product with matching SKU
		assertEqual(t, product.SKU, sku)
		assertEqual(t, product.Name, "Cardinal Jeans")

		// teardown
		repository.DeleteProduct(getDB(), sku)
	})
}

func TestGETAllProduct(t *testing.T) {
	t.Run("returns multiple products", func(t *testing.T) {
		// setup products
		prods := []repository.Product{
			repository.Product{
				SKU:      "abc125",
				Name:     "Cardinal Jeans",
				Category: "Men - Casual - Trousers",
			},
			repository.Product{
				SKU:      "abc124",
				Name:     "Cardinal T-Shirt",
				Category: "Men - Casual - T-Shirt",
			},
		}
		repository.CreateProduct(getDB(), prods[0])
		repository.CreateProduct(getDB(), prods[1])

		request, _ := http.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()

		productController := NewProductController(getDB(), log.Printf)
		productController.GetAll(response, request)

		respCode := response.Code
		expectedRespCode := 200

		assertEqual(t, respCode, expectedRespCode)

		var products []repository.Product
		err := json.NewDecoder(response.Body).Decode(&products)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%v' into Product, '%v'", response.Body, err)
		}

		// got multiple products
		assertEqual(t, len(products) > 0, true)

		// teardown
		repository.DeleteProduct(getDB(), "abc125")
		repository.DeleteProduct(getDB(), "abc124")
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("returns newly created product", func(t *testing.T) {
		sku := "abc126"
		product := repository.Product{
			SKU:      sku,
			Name:     "Cardinal Formal",
			Category: "Men - Casual - Trousers",
		}
		jsonReq, _ := json.Marshal(product)
		request, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		response := httptest.NewRecorder()

		productController := NewProductController(getDB(), log.Printf)
		productController.Create(response, request)

		respCode := response.Code
		expectedRespCode := 200

		assertEqual(t, respCode, expectedRespCode)

		err := json.NewDecoder(response.Body).Decode(&product)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into Product, '%v'", response.Body, err)
		}

		// non-zero ID means product object is persisted
		assertEqual(t, product.ID != 0, true)

		// teardown
		repository.DeleteProduct(getDB(), sku)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("returns updated product details", func(t *testing.T) {
		sku := "abc127"
		// setup product
		product := repository.Product{
			SKU:      sku,
			Name:     "Cardinal Formal",
			Category: "Men - Casual - Trousers",
		}
		repository.CreateProduct(getDB(), product)

		// update category and name
		product.Name = "Cardinal Shirt"
		product.Category = "Men - Formal - Shirt"
		jsonReq, err := json.Marshal(product)
		request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/products/%s", sku), bytes.NewBuffer(jsonReq))
		request = mux.SetURLVars(request, map[string]string{"sku": sku})
		response := httptest.NewRecorder()

		productController := NewProductController(getDB(), log.Printf)
		productController.Update(response, request)

		respCode := response.Code
		expectedRespCode := 200

		assertEqual(t, respCode, expectedRespCode)

		err = json.NewDecoder(response.Body).Decode(&product)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into Product, '%v'", response.Body, err)
		}

		// product's category is updated
		assertEqual(t, product.Name, "Cardinal Shirt")
		assertEqual(t, product.Category, "Men - Formal - Shirt")

		// teardown
		repository.DeleteProduct(getDB(), sku)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("returns response code 200", func(t *testing.T) {
		sku := "abc128"
		// setup product
		p := repository.Product{
			SKU:      sku,
			Name:     "Cardinal Jeans",
			Category: "Men - Casual - Trousers",
		}
		repository.CreateProduct(getDB(), p)

		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%s", sku), nil)
		request = mux.SetURLVars(request, map[string]string{"sku": sku})
		response := httptest.NewRecorder()

		productController := NewProductController(getDB(), log.Printf)
		productController.Delete(response, request)

		respCode := response.Code
		expectedRespCode := 204

		assertEqual(t, respCode, expectedRespCode)
	})
}

func assertEqual(t *testing.T, actual interface{}, expectation interface{}) {
	if actual != expectation {
		t.Errorf("got '%v', want '%v'", actual, expectation)
	}
}
