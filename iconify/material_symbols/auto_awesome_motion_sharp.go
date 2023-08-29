package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMotionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22V10h12v12H10Zm-4-4V6h12v2H8v10H6Zm-4-4V2h12v2H4v10H2Z"/>`),
		g.Group(children),
	)
}