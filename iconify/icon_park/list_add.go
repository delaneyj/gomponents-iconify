package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 28H24"/><path d="M8 37H24"/><path d="M8 19H40"/><path d="M8 10H40"/><path d="M30 33H40"/><path d="M35 28L35 38"/></g>`),
		g.Group(children),
	)
}