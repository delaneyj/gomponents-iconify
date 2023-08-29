package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartFourBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h20v2H2Zm1-3v-7h3v7H3Zm5 0V6h3v12H8Zm5 0V9h3v9h-3Zm5 0V3h3v15h-3Z"/>`),
		g.Group(children),
	)
}