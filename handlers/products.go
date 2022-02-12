package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/betolimasouza/go-microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products!")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}

}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) error {
	p.l.Println("ID:", id)
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Malformed Body", http.StatusBadRequest)
	}
	return data.UpdateProduct(id, prod)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Add Product!")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Malformed Body", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {
		rx := regexp.MustCompile("/([0-9]+)")
		findGroup := rx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(findGroup) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(findGroup[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := findGroup[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		err = p.updateProduct(id, rw, r)

		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		}

		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
