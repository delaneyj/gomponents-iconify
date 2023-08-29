package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 8L4 40H44L24 8Z"/><path stroke="#fff" stroke-linecap="round" d="M30 32L24 28L18 32"/><path stroke="#fff" stroke-linecap="round" d="M24 28V22"/></g>`),
		g.Group(children),
	)
}