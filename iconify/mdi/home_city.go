package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeCity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 21V10l7.5-5l7.5 5v11h-5v-7H5v7H0M24 2v19h-7V8.93l-1-.66V6h-2v.93l-4-2.66V2h14m-3 12h-2v2h2v-2m0-4h-2v2h2v-2m0-4h-2v2h2V6Z"/>`),
		g.Group(children),
	)
}