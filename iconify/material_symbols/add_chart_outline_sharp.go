package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddChartOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h2v-7H7v7Zm4 0h2V7h-2v10Zm4 0h2v-4h-2v4ZM3 21V3h11v2H5v14h14v-9h2v11H3ZM17 9V7h-2V5h2V3h2v2h2v2h-2v2h-2Zm-5 3Z"/>`),
		g.Group(children),
	)
}