package handlers

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	ghttp "github.com/maragudk/gomponents/http"
	"github.com/saent-x/fashionMNIST_Inference/core"
	"image/png"
	"net/http"
)

func HandleUpload(mux *http.ServeMux) {
	mux.Handle("/upload", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil, err
		}
		defer file.Close()

		result := Div(
			P(Class("text-xl"), Text("No Results!")),
		)
		// Make Inference
		img, err := png.Decode(file)
		if err != nil {
			result = Div(
				P(Class("text-xl"), Text("Error!")),
				P(Class("text-lg"), Text(err.Error())),
			)
		} else {
			modelPrediction, err2 := core.RunInference(img)
			if err2 != nil {
				result = Div(
					P(Class("text-xl"), Text("Error!")),
					P(Class("text-lg"), Text(err.Error())),
				)
			}

			result = Div(
				P(Class("text-2xl text-center"), Text("Model Prediction")),
				P(Class("text-xl pt-5 text-center"), Textf("[%s]", modelPrediction.ModelPrediction)),
			)

			return result, nil
		}

		return result, nil
	}))
}
