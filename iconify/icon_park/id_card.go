package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M42 8H6C4.89543 8 4 8.89543 4 10V38C4 39.1046 4.89543 40 6 40H42C43.1046 40 44 39.1046 44 38V10C44 8.89543 43.1046 8 42 8Z"/><path fill="#43CCF8" stroke="#fff" d="M36 16H28V24H36V16Z"/><path stroke="#fff" stroke-linecap="round" d="M12 32H36"/><path stroke="#fff" stroke-linecap="round" d="M12 16H18"/><path stroke="#fff" stroke-linecap="round" d="M12 24H18"/></g>`),
		g.Group(children),
	)
}