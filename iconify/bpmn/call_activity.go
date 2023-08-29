package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallActivity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<rect width="1700" height="1360" x="150" y="-627.638" fill="transparent" stroke="currentColor" stroke-linecap="round" stroke-width="220" rx="311.337" ry="306" transform="translate(0 947.638)"/>`),
		g.Group(children),
	)
}