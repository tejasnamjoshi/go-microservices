package handlers

import (
	"go-microservices/product-api-gorilla/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the list of products
// responses:
//	201: noContent

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	prod := &data.Product{}

	err = prod.DeleteProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
