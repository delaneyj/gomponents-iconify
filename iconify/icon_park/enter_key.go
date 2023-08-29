package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 44V4H24V20H4V44H44Z"/><path stroke="#fff" d="M21 28L17 32L21 36"/><path stroke="#fff" d="M34 23V32H17"/></g>`),
		g.Group(children),
	)
}