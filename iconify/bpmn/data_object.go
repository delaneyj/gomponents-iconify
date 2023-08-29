package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataObject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="90" d="M438.243 252.897h794.703l328.81 337.647v1156.564H438.244V252.897h754.486v371.647h369.028"/>`),
		g.Group(children),
	)
}