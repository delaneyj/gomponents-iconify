package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 6H4V36H13V41L23 36H44V6Z"/><path stroke="#fff" d="M14 19.5V22.5"/><path stroke="#fff" d="M24 19.5V22.5"/><path stroke="#fff" d="M34 19.5V22.5"/></g>`),
		g.Group(children),
	)
}