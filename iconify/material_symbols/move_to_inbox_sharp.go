package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveToInboxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 14l-4-4l1.4-1.45l1.6 1.6V6h2v4.15l1.6-1.6L16 10l-4 4Zm-9 7V3h18v18H3Zm9-5q.95 0 1.725-.55T14.8 14H19V5H5v9h4.2q.3.9 1.075 1.45T12 16Z"/>`),
		g.Group(children),
	)
}