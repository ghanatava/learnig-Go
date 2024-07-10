package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ghanatava/learning-Go/services/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// server MUX
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(rw, r, id)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
	p.l.Fatal(http.StatusMethodNotAllowed, " ", r.Method, " /")
}

func (p *Products) getProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println(http.StatusOK, "GET /")
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	data.AddProduct(prod)
	rw.WriteHeader(http.StatusCreated)
	p.l.Println(http.StatusCreated, "POST /")
}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request, id int) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
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
