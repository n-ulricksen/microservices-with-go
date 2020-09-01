package handlers

import (
	"net/http"
	"scratch/microservices-with-go/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
// 	200: productsResponse
// 	501: errorResponse
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
