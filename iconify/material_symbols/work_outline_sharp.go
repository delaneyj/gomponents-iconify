package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V6h6V2h8v4h6v15H2Zm2-2h16V8H4v11Zm6-13h4V4h-4v2ZM4 19V8v11Z"/>`),
		g.Group(children),
	)
}