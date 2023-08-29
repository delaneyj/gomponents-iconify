package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 11.9876L23.9938 4L32 12"/><path d="M32 36.0124L24.0061 44L16 36"/><path d="M24 4V44"/></g>`),
		g.Group(children),
	)
}