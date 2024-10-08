package handlers

import "net/http"

func Static(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
