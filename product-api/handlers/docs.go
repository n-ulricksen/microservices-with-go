// Package classification of Product API
//
// Documentation for Product API
//
// 	Schemes: http
// 	BasePath: /
// 	Version: 1.0.0
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
//
// swagger:meta
package handlers

import "scratch/microservices-with-go/product-api/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct,deleteProduct
type productIDParameterWrapper struct {
	// The id of the product being manipulated
	// in: path
	// required: true
	ID int `json:"id"`
}
