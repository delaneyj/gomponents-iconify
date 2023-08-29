package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 6H30V18H18V6Z"/><path d="M30 6H42V18H30V6Z"/><path d="M6 6H18V18H6V6Z"/><path d="M6 18H18V30H6V18Z"/><path d="M6 30H18V42H6V30Z"/></g>`),
		g.Group(children),
	)
}