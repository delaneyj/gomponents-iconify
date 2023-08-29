package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#CFD8DC" d="M35 36h4V22H26v-9h-4v9H9v14h4V26h9v10h4V26h9z"/><path fill="#3F51B5" d="M17 6h14v10H17z"/><path fill="#00BCD4" d="M32 32h10v10H32zM6 32h10v10H6zm13 0h10v10H19z"/>`),
		g.Group(children),
	)
}