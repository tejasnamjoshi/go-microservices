package handlers

import (
	"go-microservices/product-api-gorilla/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert data.", http.StatusInternalServerError)
		return
	}
}
