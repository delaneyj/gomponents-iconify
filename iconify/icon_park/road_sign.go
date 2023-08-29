package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M6 10V22H38L44 16L38 10H6Z"/><path stroke-linecap="round" d="M23 22V44"/><path stroke-linecap="round" d="M23 4V10"/><path stroke-linecap="round" d="M18 44H28"/></g>`),
		g.Group(children),
	)
}