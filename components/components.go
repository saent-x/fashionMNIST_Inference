package components

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func HomeComponent() Node {
	return Div(
		NavbarComponent(),
		InferenceComponent(),
		ResultsComponent(),
	)
}
func NavbarComponent() Node {
	return Div(Class("navbar bg-neutral"),
		Div(Class("flex-1"),
			A(Class("btn btn-ghost text-xl text-white"), Text("FashionMNIST Inference")),
		),
	)
}

func InferenceComponent() Node {
	return Div(Class("p-10"),
		Div(Class("flex flex-col"),
			Form(Attr("hx-encoding", "multipart/form-data"),
				Attr("hx-post", "/upload"),
				Attr("hx-target", "#results"),
				Input(Type("file"), Class("file-input w-full max-w-xs"), Name("file"), ID("file")),
				Button(Class("btn btn-active btn-neutral mt-10 w-[200px]"), Text("Predict"), Type("submit")),
			),
		))
}

func ResultsComponent() Node {
	return Div(Class("pt-10"),
		H2(Class("text-center text-xl"), Text("---Results---")),
		Div(Class("pt-10 w-[300px] mx-auto"), ID("results")),
	)
}
