package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 256v384H384V256h1152zm-128 256V384H512v128h896zm-128 256v384H384V768h896zm-128 256V896H512v128h640zm640 256v384H384v-384h1408zm-128 256v-128H512v128h1152zM256 1792h1664v128H128V128h128v1664z"/>`),
		g.Group(children),
	)
}