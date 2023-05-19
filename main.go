package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		password := r.URL.Query().Get("password")

		if password == "pigna" {
			fmt.Fprintf(w, "<h1 style='text-align: center'> Password corretta </h1>")
		} else {
			fmt.Fprintf(w, "<h1 style='text-align: center'> Password sbagliata </h1>")
		}
	})

	http.ListenAndServe(":8080", nil)
}
