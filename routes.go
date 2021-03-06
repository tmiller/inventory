package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tmiller/foreman-go"
	"github.com/tmiller/inventory/assets"
)

func setupRoutes() {
	http.HandleFunc("/assets/css/bootstrap.css", css_handler)
	http.HandleFunc("/servers", server_list_handler)
	http.HandleFunc("/config", config_handler)
	http.HandleFunc("/", root_handler)
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func css_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprintln(w, assets.CSS)
}

func server_list_handler(w http.ResponseWriter, r *http.Request) {

	fc := foreman.NewForemanClient(serverConfig.ForemanURL)
	fc.SetBasicAuth(
		serverConfig.ForemanUsername,
		serverConfig.ForemanPassword,
	)

	hosts := fc.Hosts()

	tmpl, _ := template.New("layoutStart").Parse(layoutStart)
	tmpl.New("layoutEnd").Parse(layoutEnd)
	tmpl.New("serverList").Parse(serverList)

	err := tmpl.ExecuteTemplate(
		w,
		"serverList",
		map[string]interface{}{
			"Title": "Servers",
			"Hosts": hosts,
		},
	)

	if err != nil {
		log.Println(err)
	}

}

func config_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, serverConfig.String())
}
