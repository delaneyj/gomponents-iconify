package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M14 13V9C14 7.89543 14.8954 7 16 7H42C43.1046 7 44 7.89543 44 9V27C44 28.1046 43.1046 29 42 29H40"/><rect width="30" height="22" x="4" y="19" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" d="M4 28L34 28"/><path stroke="#000" d="M34 23L34 35"/><path stroke="#000" d="M4 23L4 35"/><path stroke="#fff" d="M11 34L19 34"/><path stroke="#fff" d="M25 34L27 34"/></g>`),
		g.Group(children),
	)
}