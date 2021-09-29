package handlers

import (
	"go-microservices/product-api-gorilla/data"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	prod.AddProduct()
}
