package monitor

import (
	"flag"
	"net/http"
)

var addr = flag.String("Webserver Address", "localhost:8080", "Dex Equalizer Webserver")

// StartWebserver
//	Starts new web server.
func StartWebserver() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./monitor/static/index.html")
	})
	http.ListenAndServe(*addr, nil)
}
