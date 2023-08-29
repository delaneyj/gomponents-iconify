package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 14H30V26H18V14Z"/><path d="M18 26H30V38H18V26Z"/><path d="M30 14H42V26H30V14Z"/><path d="M6 26H18V38H6V26Z"/></g>`),
		g.Group(children),
	)
}