package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewWeekOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18h4V6H4v12Zm6 0h4V6h-4v12Zm6 0h4V6h-4v12Zm6 2H2V4h20v16Z"/>`),
		g.Group(children),
	)
}