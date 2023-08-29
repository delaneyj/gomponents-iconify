package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronDisable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 40H44L42 24H20C11.1634 24 4 31.1634 4 40Z"/><path stroke="#000" d="M16 8H40L42 24"/><circle cx="24" cy="24" r="9" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M26 26L24 24L22 22"/><path stroke="#fff" d="M26 22L24 24L22 26"/></g>`),
		g.Group(children),
	)
}