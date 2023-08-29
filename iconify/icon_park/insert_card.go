package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M10 18H4L4 6H44V18H38"/><path fill="#2F88FF" d="M12 12L4 41H44L36 12H12Z"/></g>`),
		g.Group(children),
	)
}