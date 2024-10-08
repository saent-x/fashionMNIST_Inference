package pages

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Page(title string, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:       title,
		Description: "-",
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/daisyui@4.12.12/dist/full.min.css"), Type("text/css")),
			Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
			Script(Src("https://cdn.tailwindcss.com")),
			Script(Attr("src", "https://unpkg.com/htmx.org@1.9.2"),
				Attr("integrity", "sha384-L6OqL9pRWyyFU3+/bjdSri+iIphTN/bvYyM37tICVyOJkWZLpP2vGn6VUEXgzg6h"),
				Attr("crossorigin", "anonymous"),
			),
		},
		Body: []Node{
			Div(Group(children)),
		},
	})
}
