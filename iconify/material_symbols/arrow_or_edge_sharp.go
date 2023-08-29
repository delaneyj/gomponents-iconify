package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOrEdgeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.5 21l-4-4l1.4-1.4l1.6 1.575V11H1V3h2v6h5.5v8.175l1.575-1.575L11.5 17l-4 4Zm9 0l-4-4l1.4-1.4l1.6 1.575V9H21V3h2v8h-5.5v6.2l1.575-1.6L20.5 17l-4 4Z"/>`),
		g.Group(children),
	)
}