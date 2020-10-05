package main

import (
	"net/http"

	"github.com/msoerjanto/pikaso/data"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	artists, err := data.Artists()
	if err != nil {
		error_message(writer, request, "Cannot get artists")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, artists, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, artists, "layout", "private.navbar", "index")
		}
	}
}
