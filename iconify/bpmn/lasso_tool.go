package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LassoTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linejoin="round" stroke-width="205.238" d="M1304.97 699.019H868.856"/><path stroke-width="80" d="M1566.732 696.368h285.2v273.246m.001 592.034v273.247h-277.985m277.304-652.426v153.246m-1140.123 228.21v273.247h277.984m209.817 0h165.201"/><path stroke-linejoin="round" stroke-width="205.238" d="M708.49 104.8v436.115m0 323.815v436.114M545.042 699.019H108.927"/></g>`),
		g.Group(children),
	)
}