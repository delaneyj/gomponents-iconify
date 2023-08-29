package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnknownMirroredSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0Q828 0 705 34t-229 97t-194 150t-150 194t-97 230T0 960q0 132 34 255t97 229t150 194t194 150t230 97t255 35q132 0 255-34t229-97t194-150t150-194t97-229t35-256q0-132-34-255t-97-229t-150-194t-194-150t-229-97T960 0zm-64 1408h128v128H896v-128zm-8-169q0-37-7-70t-36-62q-39-39-77-74t-68-75t-49-85t-19-105q0-68 26-127t70-104t104-71t128-26q68 0 127 26t104 70t71 104t26 128h-144q0-38-14-71t-40-59t-58-39t-72-15q-38 0-71 14t-59 40t-39 58t-15 72q0 41 19 73t47 61t62 60t61 66t48 81t19 107v64H888v-41z"/>`),
		g.Group(children),
	)
}