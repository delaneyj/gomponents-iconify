package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 18H30V30H18V18Z"/><path d="M30 18H42V30H30V18Z"/><path d="M6 18H18V30H6V18Z"/><path d="M18 30H30V42H18V30Z"/><path d="M18 6H30V18H18V6Z"/></g>`),
		g.Group(children),
	)
}