package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 22v-1.5l2-1.5v-5.5L2 16v-2l8.5-5V3.5q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5V9l8.5 5v2l-8.5-2.5V19l2 1.5V22L12 21l-3.5 1Z"/>`),
		g.Group(children),
	)
}