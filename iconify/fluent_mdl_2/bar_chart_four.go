package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 2048v-640h384v640H0zm512 0V896h384v1152H512zM1536 896h384v1152h-384V896zm-512 1152V384h384v1664h-384z"/>`),
		g.Group(children),
	)
}