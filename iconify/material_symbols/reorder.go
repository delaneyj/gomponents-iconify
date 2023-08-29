package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19v-2h18v2H3Zm0-4v-2h18v2H3Zm0-4V9h18v2H3Zm0-4V5h18v2H3Z"/>`),
		g.Group(children),
	)
}