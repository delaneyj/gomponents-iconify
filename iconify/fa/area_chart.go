package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M2048 1408v128H0V0h128v1408h1920zM1664 384l256 896H256V704l448-576l576 576z"/>`),
		g.Group(children),
	)
}