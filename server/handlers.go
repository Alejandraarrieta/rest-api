package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	m"rest-api/models"
)
var products[]*m.Product = []*m.Product{}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed\n")
		return
	}
	fmt.Fprintf(w, "Hello there %s\n", "visitor")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func addProduct(w http.ResponseWriter, r *http.Request) {
	product := &m.Product{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	products = append(products, product)
	fmt.Fprintf(w, "products was added")

}
