package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tmiller/fapi"
	"github.com/tmiller/inventory/assets"
)

var serverSettings settings

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/servers'>Server List</a>")
}

func css_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprintln(w, assets.CSS)
}

func server_list_handler(w http.ResponseWriter, r *http.Request) {

	fc := fapi.NewForemanClient(serverSettings.ForemanURL)
	fc.SetBasicAuth(
		serverSettings.ForemanUsername,
		serverSettings.ForemanPassword,
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

func settings_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, serverSettings.String())
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	serverSettings = loadSettings("settings.json")
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/assets/css/bootstrap.css", css_handler)
	http.HandleFunc("/servers", server_list_handler)
	http.HandleFunc("/settings", settings_handler)

	address := fmt.Sprintf(":%s", serverSettings.ServerPort)
	http.ListenAndServe(address, nil)
}
