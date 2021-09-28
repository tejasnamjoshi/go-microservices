package handlers

import (
	"go-microservices/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert data.", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to convert data.", http.StatusBadRequest)
		return
	}

	prod.AddProduct()
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = prod.UpdateProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

// improved
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
		p.l.Println(r.URL.Path)
		re := regexp.MustCompile(`/([0-9]+)`)
		path := r.URL.Path

		m := re.FindAllStringSubmatch(path, -1)
		if len(m) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(m[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := m[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}

		p.l.Println("Got ID", id)
		p.updateProduct(id, rw, r)

		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// simple
// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	p.l.Println("Products")
// 	lp := data.GetProducts()
// 	d, err := json.Marshal(lp)
// 	if err != nil {
// 		http.Error(rw, "Unable to convert data.", http.StatusInternalServerError)
// 		return
// 	}

// 	rw.Write(d)
// }
