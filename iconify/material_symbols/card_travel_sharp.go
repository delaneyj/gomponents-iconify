package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTravelSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V6h5V2h10v4h5v15H2ZM9 6h6V4H9v2ZM4 17h16v-3H4v3Z"/>`),
		g.Group(children),
	)
}