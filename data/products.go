package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"-"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"LastUpdate"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	index, err := findProduct(id)

	if err != nil {
		return err
	}

	productList[index] = p
	productList[index].UpdateOn = time.Now().Local().String()

	return err
}

var ErrProductNotFound = fmt.Errorf("Product not Found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}

	return -1, ErrProductNotFound
}

func GetNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Café",
		Description: "Café simples",
		Price:       2.45,
		SKU:         "abc234",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Expresso",
		Description: "Café italiano",
		Price:       1.99,
		SKU:         "fff234",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
