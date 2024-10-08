package pages

import (
	. "github.com/maragudk/gomponents"
	ghttp "github.com/maragudk/gomponents/http"
	"github.com/saent-x/fashionMNIST_Inference/components"
	"net/http"
)

func Home(mux *http.ServeMux) {
	mux.Handle("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return Page("Home", components.HomeComponent()), nil
	}))
}
