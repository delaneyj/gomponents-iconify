package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnknownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0q132 0 255 34t229 97t194 150t150 194t97 230t35 255q0 132-34 255t-97 229t-150 194t-194 150t-230 97t-255 35q-132 0-255-34t-229-97t-194-150t-150-194t-97-229T0 960q0-132 34-255t97-229t150-194t194-150t229-97T960 0zm64 1408H896v128h128v-128zm8-169q0-37 7-70t36-62q39-39 77-74t68-75t49-85t19-105q0-68-26-127t-70-104t-104-71t-128-26q-68 0-127 26t-104 70t-71 104t-26 128h144q0-38 14-71t40-59t58-39t72-15q38 0 71 14t59 40t39 58t15 72q0 41-19 73t-47 61t-62 60t-61 66t-48 81t-19 107v64h144v-41z"/>`),
		g.Group(children),
	)
}