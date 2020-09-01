package handlers

import (
	"log"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
