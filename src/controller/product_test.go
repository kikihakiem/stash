package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kikihakiem/stash/go/simple-crud/repository"
)

func TestGETSingleProduct(t *testing.T) {
	t.Run("returns product details", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/products/abc123", nil)
		response := httptest.NewRecorder()

		ProductController(response, request)

		respCode := response.Code
		expectedRespCode := 200

		assertEqual(t, respCode, expectedRespCode)

		var product repository.Product
		err := json.NewDecoder(response.Body).Decode(&product)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into Product, '%v'", response.Body, err)
		}

		assertEqual(t, product.SKU, "abc123")
	})
}

func assertEqual(t *testing.T, actual interface{}, expectation interface{}) {
	if actual != expectation {
		t.Errorf("got '%v', want '%v'", actual, expectation)
	}
}
