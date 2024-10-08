package handlers

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	ghttp "github.com/maragudk/gomponents/http"
	"net/http"
)

func HandleUpload(mux *http.ServeMux) {
	mux.Handle("/upload", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil, err
		}
		defer file.Close()

		result := Div(
			P(Class("text-lg"), Textf("Uploaded file: %s", header.Filename)),
			P(Class("text-lg"), Textf("File size: %d bytes", header.Size)),
		)

		return result, nil
	}))

}
