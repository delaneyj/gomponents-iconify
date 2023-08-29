package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 512v1152H384V512h384zM640 1536V640H512v896h128zm640-768v896H896V768h384zm-128 768V896h-128v640h128zm640-1280v1408h-384V256h384zm-128 1280V384h-128v1152h128zM256 1792h1664v128H128V128h128v1664z"/>`),
		g.Group(children),
	)
}