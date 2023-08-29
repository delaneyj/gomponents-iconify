package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedColumnChartTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 256v1408h-384V256h384zm-128 384V384h-128v256h128zm-384 128v896H896V768h384zm-128 384V896h-128v256h128zM768 512v1152H384V512h384zM640 896V640H512v256h128zm-384 896h1664v128H128V128h128v1664z"/>`),
		g.Group(children),
	)
}