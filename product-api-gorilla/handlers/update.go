package handlers

import (
	"go-microservices/product-api-gorilla/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = prod.UpdateProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
