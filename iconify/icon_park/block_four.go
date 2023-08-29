package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 12H18V24H6V12Z"/><path d="M18 12H30V24H18V12Z"/><path d="M30 12H42V24H30V12Z"/><path d="M18 24H30V36H18V24Z"/></g>`),
		g.Group(children),
	)
}