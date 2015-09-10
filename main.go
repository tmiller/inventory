package main

import (
	"fmt"
	"net/http"
)

var serverSettings settings

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func server_list_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func settings_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, serverSettings.String())
}

func main() {

	serverSettings = loadSettings("foo")
	http.HandleFunc("/", handler)
	http.HandleFunc("/servers", server_list_handler)
	http.HandleFunc("/settings", settings_handler)

	address := fmt.Sprintf(":%s", serverSettings.ServerPort)
	http.ListenAndServe(address, nil)
}
