// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"go-microservices/product-api-gorilla/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// A list of products that are returned in response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the system
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productNoContent struct{}

type KeyProduct struct{}

func (prod Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
