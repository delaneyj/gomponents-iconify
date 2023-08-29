package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusCircleQuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1536h128v128H896v-128zm64-960q66 0 124 25t101 69t69 102t26 124q0 60-19 104t-47 81t-62 65t-61 59t-48 63t-19 76v64H896v-64q0-60 19-104t47-81t62-65t61-59t48-63t19-76q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75H640q0-66 25-124t68-101t102-69t125-26z"/>`),
		g.Group(children),
	)
}