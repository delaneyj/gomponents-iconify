package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="34" cy="14" r="9"/><circle cx="12" cy="25" r="7"/><circle cx="29" cy="37" r="5"/></g>`),
		g.Group(children),
	)
}