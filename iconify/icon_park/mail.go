package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 39H44V24V9H24H4V24V39Z"/><path stroke="#fff" stroke-linecap="round" d="M4 9L24 24L44 9"/><path stroke="#000" stroke-linecap="round" d="M24 9H4V24"/><path stroke="#000" stroke-linecap="round" d="M44 24V9H24"/></g>`),
		g.Group(children),
	)
}