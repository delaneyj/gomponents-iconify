package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-dashoffset="200" stroke-width="40" d="M799.876 406.976v159.072H587.842v119.547h212.034v159.073l232.942-218.846z"/><path stroke-width="90" d="M438.243 252.897h794.702l328.812 337.647v1156.564H438.243V252.897h754.486v371.647h369.028"/></g>`),
		g.Group(children),
	)
}