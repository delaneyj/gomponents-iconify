package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMmaSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-4h10v4H7Zm-1-5l-1-5.05V3h12v4h2v3.95L18 16H6Zm2-6h6V7H8v3Z"/>`),
		g.Group(children),
	)
}