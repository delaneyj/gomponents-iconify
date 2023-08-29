package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Log(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="28" height="34" x="13" y="10" fill="#2F88FF" stroke="#000"/><path stroke="#000" stroke-linecap="round" d="M35 10V4H8C7.44772 4 7 4.44772 7 5V38H13"/><path stroke="#fff" stroke-linecap="round" d="M21 22H33"/><path stroke="#fff" stroke-linecap="round" d="M21 30H33"/></g>`),
		g.Group(children),
	)
}