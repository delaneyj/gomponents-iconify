package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 10C4 8.89543 4.89543 8 6 8H42C43.1046 8 44 8.89543 44 10V38C44 39.1046 43.1046 40 42 40H6C4.89543 40 4 39.1046 4 38V10Z"/><path stroke="#fff" stroke-linecap="square" d="M4 16H44"/><path stroke="#fff" stroke-linecap="round" d="M27 32H36"/><path stroke="#000" stroke-linecap="round" d="M44 10V26"/><path stroke="#000" stroke-linecap="round" d="M4 10V26"/></g>`),
		g.Group(children),
	)
}