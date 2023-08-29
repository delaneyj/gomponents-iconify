package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalSplitOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7V5h18v2Zm0 4V9h18v2Zm0 8v-6h18v6Zm2-2h14v-2H5Zm0 0v-2v2Z"/>`),
		g.Group(children),
	)
}