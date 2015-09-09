package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func server_list_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/servers", server_list_handler)
	http.ListenAndServe(":8080", nil)
}
