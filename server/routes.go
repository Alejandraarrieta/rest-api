package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			getProduct(w, r)

		case http.MethodPost:
			addProduct(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})
}
