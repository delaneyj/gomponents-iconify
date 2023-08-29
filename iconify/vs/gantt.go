package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gantt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M198 992h154V480H70q-28 0-49-21T0 410V166q0-28 21-49t49-21h282V32q0-14 9-23t23-9t23 9t9 23v64h448V32q0-14 9-23t23-9t23 9t9 23v64h154q29 0 49.5 20.5T1152 166v244q0 28-21 49t-49 21H928v64h448V32q0-14 9-23t23-9t23 9t9 23v512h282q29 0 49.5 20.5T1792 614v244q0 28-21 49t-49 21h-282v512q0 14-9 23t-23 9t-23-9t-9-23V928H928v64h282q28 0 49 21t21 49v244q0 28-21 49t-49 21H928v64q0 14-9 23t-23 9t-23-9t-9-23v-64H416v64q0 14-9 23t-23 9t-23-9t-9-23v-64H198q-28 0-49-21t-21-49v-244q0-28 21-49t49-21zm666 0v-64H710q-28 0-49-21t-21-49V614q0-29 20.5-49.5T710 544h154v-64H416v512h448z"/>`),
		g.Group(children),
	)
}