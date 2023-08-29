package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 40H44L42 24H20C11.1634 24 4 31.1634 4 40Z"/><path stroke="#000" d="M16 8H40L42 24"/><path stroke="#fff" d="M17 32H19"/></g>`),
		g.Group(children),
	)
}