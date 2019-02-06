package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kikihakiem/stash/go/simple-crud/repository"
)

func TestGETSingleProduct(t *testing.T) {
	t.Run("returns product details", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/products/abc123", nil)
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

func TestCreateProduct(t *testing.T) {
	t.Run("returns product details", func(t *testing.T) {
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

func assertEqual(t *testing.T, actual interface{}, expectation interface{}) {
	if actual != expectation {
		t.Errorf("got '%v', want '%v'", actual, expectation)
	}
}
