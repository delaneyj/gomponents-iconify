package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedbackRequestSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1408H704l-448 448v-448H0V128h2048zM1024 1280H896v128h128v-128zm8-169q0-37 7-70t36-62q39-39 77-74t68-75t49-85t19-105q0-68-26-127t-70-104t-104-71t-128-26q-68 0-127 26t-104 70t-71 104t-26 128h144q0-38 14-71t40-59t58-39t72-15q38 0 71 14t59 40t39 58t15 72q0 41-19 73t-47 61t-62 60t-61 66t-48 81t-19 107v64h144v-41z"/>`),
		g.Group(children),
	)
}