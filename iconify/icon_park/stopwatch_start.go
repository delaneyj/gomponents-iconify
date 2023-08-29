package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C33.3888 44 41 36.3888 41 27C41 17.6112 33.3888 10 24 10C14.6112 10 7 17.6112 7 27C7 36.3888 14.6112 44 24 44Z"/><path stroke="#000" stroke-linecap="round" d="M18 4H30"/><path stroke="#fff" stroke-linecap="round" d="M24 19V27"/><path stroke="#fff" stroke-linecap="round" d="M32 27H24"/><path stroke="#000" stroke-linecap="round" d="M24 4V8"/></g>`),
		g.Group(children),
	)
}