package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rowing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-miterlimit="2" stroke-width="4"><path fill="#2F88FF" d="M30.0195 16C32.781 16 35.0195 13.7614 35.0195 11C35.0195 8.23858 32.781 6 30.0195 6C27.2581 6 25.0195 8.23858 25.0195 11C25.0195 13.7614 27.2581 16 30.0195 16Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 31L35 44"/><path stroke-linecap="round" stroke-linejoin="round" d="M32.01 40L26 29L9 40L21 20L44 25"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 40H44"/></g>`),
		g.Group(children),
	)
}