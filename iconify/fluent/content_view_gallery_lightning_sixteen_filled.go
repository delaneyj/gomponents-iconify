package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryLightningSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2H4.5A2.5 2.5 0 0 0 2 4.5v7A2.5 2.5 0 0 0 4.5 14h4.22l.25-1h-.468a1.497 1.497 0 0 1-1.415-1H4.5a.5.5 0 0 1 0-1h2.587l.03-.077l.385-.923H4.5a.5.5 0 0 1 0-1h3c.14 0 .265.057.356.149L9.2 5.923a1.5 1.5 0 0 1 .8-.804V2Zm4 2.5V5h-3V2h.5A2.5 2.5 0 0 1 14 4.5Zm-10 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-3ZM5 5v2h2V5H5Zm3.502 7h1.75l-.59 2.36c-.12.482.462.826.826.486l4.873-4.548A.75.75 0 0 0 14.849 9h-1.097l.78-2.342A.5.5 0 0 0 14.059 6h-3.473a.5.5 0 0 0-.461.308l-2.084 5a.5.5 0 0 0 .462.692Z"/>`),
		g.Group(children),
	)
}