package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="14" height="9" x="17" y="6" fill="#2F88FF"/><rect width="14" height="9" x="6" y="33" fill="#2F88FF"/><rect width="14" height="9" x="28" y="33" fill="#2F88FF"/><path stroke-linecap="round" d="M24 16V24"/><path stroke-linecap="round" d="M13 33V24H35V33"/></g>`),
		g.Group(children),
	)
}