package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 4H4V44H44V4Z"/><path stroke="#fff" stroke-linecap="round" d="M16 4V16H4"/><path stroke="#fff" stroke-linecap="round" d="M36 24V36H24"/><path stroke="#fff" stroke-linecap="round" d="M36 36L24 24"/><path stroke="#000" stroke-linecap="round" d="M4 6V26"/><path stroke="#000" stroke-linecap="round" d="M7 4H27"/></g>`),
		g.Group(children),
	)
}