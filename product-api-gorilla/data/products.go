package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
//swagger:model
type Product struct {
	// the id for this user
	//
	// required: true
	//min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

type Products []*Product

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Expresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func (p Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) AddProduct() {
	p.ID = getNextID()
	productList = append(productList, p)
}

func (p *Product) UpdateProduct(id int) error {
	pos, err := getProductPosById(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func (p *Product) DeleteProduct(id int) error {
	pos, err := getProductPosById(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)

	return nil
}

func getProductPosById(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}

	return -1, errors.New("Product Not Found")
}

func getNextID() int {
	return len(productList) + 1
}
