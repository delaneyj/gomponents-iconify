package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Text(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6v2H3V6h18M3 18h9v-2H3v2m0-5h18v-2H3v2Z"/>`),
		g.Group(children),
	)
}