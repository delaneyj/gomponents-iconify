package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 30H18V42H6V30Z"/><path d="M18 18H30V30H18V18Z"/><path d="M30 6H42V18H30V6Z"/></g>`),
		g.Group(children),
	)
}