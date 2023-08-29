package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartVerticalFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 384v768h-128V256h384v896h-128V384h-128zm-384 1152l128 128H768V768h384v384h-128V896H896v640h128zM640 512v1152H256V512h384zM512 1536V640H384v896h128zm640-256h896v152l-256 288v328h-384v-328l-256-288v-152zm512 392l234-264h-596l234 264v248h128v-248zM128 128v1664h1152v128H0V128h128z"/>`),
		g.Group(children),
	)
}