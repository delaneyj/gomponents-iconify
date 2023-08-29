package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WristOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21.9L6.1 17H1v-2h5.9l3.45 3.425L8.875 15.5H21v2h-8.875l1.125 2.225L11 21.9ZM1 9V7h4.6l1.95-2H20v2H8.425L6.4 9H1Zm13 5v-2h9v2h-9Zm0-3.5v-2h8v2h-8ZM1 12.7Z"/>`),
		g.Group(children),
	)
}