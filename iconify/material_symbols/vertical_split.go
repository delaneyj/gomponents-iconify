package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13v2h8v-2H3Zm0 4v2h8v-2H3Zm0-8v2h8V9H3Zm0-4v2h8V5H3Zm10 0h8v14h-8V5Z"/>`),
		g.Group(children),
	)
}