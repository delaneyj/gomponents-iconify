package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShadowMinusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V6h4V2h16v16h-4v4H2Zm6-6h12V4H8v12Zm2-5V9h8v2h-8Z"/>`),
		g.Group(children),
	)
}