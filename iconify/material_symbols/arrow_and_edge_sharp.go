package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAndEdgeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-4-4l1.4-1.4l1.6 1.575V13q0-.825-.588-1.413T9 11H1V3h2v6h6q.9 0 1.675.363T12 10.35q.55-.625 1.325-.987T15 9h6V3h2v8h-8q-.825 0-1.413.588T13 13v4.175l1.575-1.575L16 17l-4 4Z"/>`),
		g.Group(children),
	)
}