package main

import (
	"go-microservices/product-api-gorilla/client/client"
	"go-microservices/product-api-gorilla/client/client/products"
	"testing"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()

	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(prod.GetPayload()) == 0 {
		t.Fatal(err)
	}
}
