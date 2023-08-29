package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiFoodBeverageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2h16v2H4ZM18 8h2V5h-2v3ZM4 17V3h5v2.4L7 7v5h5V7l-2-1.6V3h10q.825 0 1.413.588T22 5v3q0 .825-.588 1.413T20 10h-2v7H4Z"/>`),
		g.Group(children),
	)
}