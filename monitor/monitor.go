package monitor

import (
	"net/http"
)

// SetWebHandler
//	Starts new web server.
func SetWebHandler() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./monitor/static/index.html")
	})
}
