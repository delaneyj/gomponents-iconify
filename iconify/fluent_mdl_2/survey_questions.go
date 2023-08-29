package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveyQuestions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1663h128v128H896v-128zm64-960q66 0 124 25t101 69t69 102t26 124q0 60-19 105t-47 80t-62 65t-61 60t-48 63t-19 75v64H896v-64q0-60 19-104t47-80t62-66t61-59t48-63t19-76q0-39-15-74t-41-61t-62-42t-74-15q-39 0-74 15t-61 41t-42 62t-15 74H640q0-66 25-124t68-101t102-69t125-26zm832-154v1499H128V0h1120l544 549zm-507-37h290l-290-292v292zm379 1408V640h-507V128H256v1792h1408z"/>`),
		g.Group(children),
	)
}