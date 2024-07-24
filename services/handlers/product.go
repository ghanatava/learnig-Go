// Package classification of Product API.
//
// Documentation for Product API.
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ghanatava/learning-Go/services/data"
	"github.com/gorilla/mux"
)

// Products represents the handlers for the product API
type Products struct {
	l *log.Logger
}

// NewProduct creates a new product handler
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts handles the GET /products endpoint
// swagger:route GET /products products getProducts
//
// Responses:
//
//	200: productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println(http.StatusOK, "GET /products")
}

// AddProducts handles the POST /products endpoint
// swagger:route POST /products products addProduct
//
// Responses:
//
//	201: productCreatedResponse
//	400: badRequestResponse
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	rw.WriteHeader(http.StatusCreated)
	p.l.Println(http.StatusCreated, "POST /products")
}

// UpdateProducts handles the PUT /products/{id} endpoint
// swagger:route PUT /products/{id} products updateProduct
//
// Responses:
//
//	204: noContentResponse
//	400: badRequestResponse
//	404: notFoundResponse
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.l.Printf("%v %v %v\n", http.StatusNotFound, r.Method, r.URL)
		return
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		p.l.Printf("%v %v %v\n", http.StatusInternalServerError, r.Method, r.URL)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
	p.l.Println(http.StatusNoContent, "PUT /products/{id}")
}

// KeyProduct is a key used for product context
type KeyProduct struct{}

// MiddlewareProductValidation validates the product in the request
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

// Product response types for Swagger

// swagger:response productsResponse
type productsResponse struct {
	// in: body
	Body []data.Product
}

// swagger:response productCreatedResponse
type productCreatedResponse struct {
}

// swagger:response badRequestResponse
type badRequestResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response notFoundResponse
type notFoundResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response noContentResponse
type noContentResponse struct {
}
