package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21v-7q0-.825.588-1.413T7 12h1v-2H7q-.825 0-1.413-.588T5 8V5q0-.825.588-1.413T7 3h10q.825 0 1.413.588T19 5v3q0 .825-.588 1.413T17 10h-1v2h1q.825 0 1.413.588T19 14v7h-2v-3H7v3H5Zm5-9h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}