package http

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello the local 8080!")

}
func startServer() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)

}
