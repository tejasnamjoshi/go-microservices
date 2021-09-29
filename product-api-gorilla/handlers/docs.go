package handlers

import "go-microservices/product-api-gorilla/data"

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
