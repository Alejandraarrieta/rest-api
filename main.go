package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed\n")
			return
		}
		fmt.Fprintf(w, "Hello there %s\n", "visitor")
	})

	srv := http.Server{
		Addr: ":9000",
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
