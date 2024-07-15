package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ghanatava/learning-Go/services/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println(http.StatusOK, "GET /")
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	data.AddProduct(prod)
	rw.WriteHeader(http.StatusCreated)
	p.l.Println(http.StatusCreated, "POST /")
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	//Get Id from url itself

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to covert id", http.StatusBadRequest)
	}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.l.Printf("%v %v %v\n", http.StatusNotFound, r.Method, r.URL)
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		p.l.Printf("%v %v %v\n", http.StatusInternalServerError, r.Method, r.URL)
	}
}
