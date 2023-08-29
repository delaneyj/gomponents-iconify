package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 768v896H896V768h384zm512-512v1408h-384V256h384zM256 1792h1664v128H128V128h128v1664zM768 512v1152H384V512h384z"/>`),
		g.Group(children),
	)
}