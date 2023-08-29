package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14v2H2V2h14v4h-2V4H4v10h2Zm2 8V8h14v14H8Z"/>`),
		g.Group(children),
	)
}