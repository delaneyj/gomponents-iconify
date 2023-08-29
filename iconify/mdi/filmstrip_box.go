package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2V5c0-1.1.9-2 2-2h14M7 18v-2H5v2h2m0-5v-2H5v2h2m0-5V6H5v2h2m12 10v-2h-2v2h2m0-5v-2h-2v2h2m0-5V6h-2v2h2Z"/>`),
		g.Group(children),
	)
}