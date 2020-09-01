package handlers

import (
	"net/http"
	"scratch/microservices-with-go/product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a product's details
//
// responses:
// 	201: noContentResponse
// 	404: errorResponse
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.logger.Println("Handle PUT Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
