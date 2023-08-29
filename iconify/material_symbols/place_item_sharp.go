package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceItemSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V7h6v2H5v10h14V9h-4V7h6v14H3Zm9-5l-4-4l1.4-1.4l1.6 1.575V0h2v12.175l1.6-1.575L16 12l-4 4Z"/>`),
		g.Group(children),
	)
}