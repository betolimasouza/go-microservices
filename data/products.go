package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"-"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func GetProducts() []*Product {
	return productList
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
