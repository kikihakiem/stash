package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/kikihakiem/stash/go/simple-crud/repository"
)

func TestGETSingleProduct(t *testing.T) {
	t.Run("returns product details", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/products/abc123", nil)
		request = mux.SetURLVars(request, map[string]string{"sku": "abc123"})
		response := httptest.NewRecorder()

		productController := NewProductController(nil, log.Printf)
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
		assertEqual(t, product.SKU, "abc123")
	})
}

func TestGETAllProduct(t *testing.T) {
	t.Run("returns multiple products", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()

		productController := NewProductController(nil, log.Printf)
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
		assertEqual(t, len(products), 2)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("returns newly created product", func(t *testing.T) {
		product := repository.Product{
			SKU:      "foo345",
			Name:     "Cardinal Formal",
			Category: "Men - Casual - Trousers",
		}
		jsonReq, _ := json.Marshal(product)
		request, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		response := httptest.NewRecorder()

		productController := NewProductController(nil, log.Printf)
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
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("returns updated product details", func(t *testing.T) {
		product := repository.Product{
			SKU:      "foo345",
			Name:     "Cardinal Formal",
			Category: "Men - Casual - Trousers",
		}
		jsonReq, err := json.Marshal(product)
		request, err := http.NewRequest(http.MethodPatch, "/products", bytes.NewBuffer(jsonReq))
		request = mux.SetURLVars(request, map[string]string{"sku": "foo345"})
		response := httptest.NewRecorder()

		productController := NewProductController(nil, log.Printf)
		productController.Update(response, request)

		respCode := response.Code
		expectedRespCode := 200

		assertEqual(t, respCode, expectedRespCode)

		err = json.NewDecoder(response.Body).Decode(&product)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into Product, '%v'", response.Body, err)
		}

		// product's category is updated
		assertEqual(t, product.Category, "Men - Casual - Pants")
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("returns response code 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodDelete, "/products/abc123", nil)
		request = mux.SetURLVars(request, map[string]string{"sku": "abc123"})
		response := httptest.NewRecorder()

		productController := NewProductController(nil, log.Printf)
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
