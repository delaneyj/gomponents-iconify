package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOrEdge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.5 21l-4-4l1.4-1.4l1.6 1.575V11H3q-.825 0-1.413-.588T1 9V3h2v6h3.5q.825 0 1.413.588T8.5 11v6.175l1.575-1.575L11.5 17l-4 4Zm9 0l-4-4l1.4-1.4l1.6 1.575V11q0-.825.588-1.413T17.5 9H21V3h2v6q0 .825-.588 1.413T21 11h-3.5v6.2l1.575-1.6L20.5 17l-4 4Z"/>`),
		g.Group(children),
	)
}