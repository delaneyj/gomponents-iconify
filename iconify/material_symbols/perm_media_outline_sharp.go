package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermMediaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V6h2v13h17v2H1Zm4-4V2h7l2 2h9v13H5Zm2-2h14V6h-7.825l-2-2H7v11Zm0 0V4v11Zm2-2h10l-3.45-4.5l-2.3 3l-1.55-2L9 13Z"/>`),
		g.Group(children),
	)
}