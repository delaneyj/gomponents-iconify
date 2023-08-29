package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WristSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21.9L6.1 17H1V7h4.6l1.95-2H20v2h-6v1.5h8v2h-8V12h9v2h-9v1.5h7v2h-8.875l1.125 2.225L11 21.9Z"/>`),
		g.Group(children),
	)
}