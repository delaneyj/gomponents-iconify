package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viewfinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 6H8C6.89543 6 6 6.89543 6 8V16"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 42H8C6.89543 42 6 41.1046 6 40V32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M32 42H40C41.1046 42 42 41.1046 42 40V32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M32 6H40C41.1046 6 42 6.89543 42 8V16"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M13 33V18H19L21 15H27L29 18H35V33H13Z"/><path fill="#43CCF8" stroke="#fff" stroke-miterlimit="10" d="M24 28C25.6569 28 27 26.6569 27 25C27 23.3431 25.6569 22 24 22C22.3431 22 21 23.3431 21 25C21 26.6569 22.3431 28 24 28Z"/></g>`),
		g.Group(children),
	)
}