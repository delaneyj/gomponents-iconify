package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedColumnChartTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1792v128H128V128h128v1664h1664zM768 1664H384V512h384v1152zm-128-640H512v512h128v-512zm0-384H512v256h128V640zm640 1024H896V768h384v896zm-128-384h-128v256h128v-256zm0-384h-128v256h128V896zm640 768h-384V256h384v1408zm-128-896h-128v768h128V768zm0-384h-128v256h128V384z"/>`),
		g.Group(children),
	)
}