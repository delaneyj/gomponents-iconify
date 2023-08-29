package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeDiagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="10" cy="24" r="4" fill="#2F88FF"/><circle cx="38" cy="10" r="4" fill="#2F88FF"/><circle cx="38" cy="24" r="4" fill="#2F88FF"/><circle cx="38" cy="38" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 38L22 38V10H34"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 24L34 24"/></g>`),
		g.Group(children),
	)
}