package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tmiller/fapi"
)

var serverSettings settings

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func server_list_handler(w http.ResponseWriter, r *http.Request) {

	fc := fapi.NewForemanClient(serverSettings.ForemanURL)
	fc.SetBasicAuth(
		serverSettings.ForemanUsername,
		serverSettings.ForemanPassword,
	)

	hosts := fc.Hosts()

	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "<ul>")
	for _, host := range hosts {
		fmt.Fprintf(w, "<li>%s</li>\n", host.Name)
	}
	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "</body></html>")
}

func settings_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, serverSettings.String())
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	serverSettings = loadSettings("settings.json")
	http.HandleFunc("/", handler)
	http.HandleFunc("/servers", server_list_handler)
	http.HandleFunc("/settings", settings_handler)

	address := fmt.Sprintf(":%s", serverSettings.ServerPort)
	http.ListenAndServe(address, nil)
}
