package handlers

import (
	"net/http"
	"scratch/microservices-with-go/product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the data store
// responses:
// 	201: noContentResponse
// 	404: errorResponse
// 	501: errorResponse

// DeleteProducts deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.logger.Println("Handle DELETE Product", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
