package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiFoodBeverageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2h16v2H4Zm0-4V3h16q.825 0 1.413.588T22 5v3q0 .825-.588 1.413T20 10h-2v7H4ZM6 5h10H6Zm12 3h2V5h-2v3Zm-2 7V5h-6v.4L12 7v5H7V7l2-1.6V5H6v10h10ZM9 5h1h-1Z"/>`),
		g.Group(children),
	)
}