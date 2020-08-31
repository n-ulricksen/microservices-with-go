package handlers

import (
	"log"
	"net/http"
	"regexp"
	"scratch/microservices-with-go/product-api/data"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		match := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(match) != 1 {
			p.logger.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		if len(match[0]) != 2 {
			p.logger.Println("Invalid URI more than capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		idString := match[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.logger.Println("Invalid URI unable to convert to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		p.updateProducts(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET products")

	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
