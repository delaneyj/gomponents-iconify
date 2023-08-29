package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeBringToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v4H9V4H4v5h2v2H2V2m20 11v9h-9v-4h2v2h5v-5h-2v-2h4M8 8h8v8H8V8Z"/>`),
		g.Group(children),
	)
}