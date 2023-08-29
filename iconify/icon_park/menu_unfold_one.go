package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfoldOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 10.5H40"/><path stroke-linecap="round" d="M24 19.5H40"/><path stroke-linecap="round" d="M24 28.5H40"/><path stroke-linecap="round" d="M8 37.5H40"/><path fill="#2F88FF" d="M16 19L8 24L16 29V19Z"/></g>`),
		g.Group(children),
	)
}