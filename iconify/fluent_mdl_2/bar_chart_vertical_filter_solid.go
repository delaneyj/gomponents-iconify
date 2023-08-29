package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartVerticalFilterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 1536l128 128H768V768h384v384h-128v384zm640-1280v896h-384V256h384zM640 512v1152H256V512h384zM128 128v1664h1152v128H0V128h128zm1024 1152h896v152l-256 288v328h-384v-328l-256-288v-152z"/>`),
		g.Group(children),
	)
}