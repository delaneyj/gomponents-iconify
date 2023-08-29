package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M44 18V8H4V18"/><path fill="#2F88FF" stroke="#000" d="M44 18H4V40H44V18Z"/><path stroke="#fff" stroke-linecap="round" d="M12 29H14"/><path stroke="#fff" stroke-linecap="round" d="M20 29H22"/><path stroke="#fff" stroke-linecap="round" d="M28 29H30"/></g>`),
		g.Group(children),
	)
}