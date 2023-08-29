package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M21 38C30.3888 38 38 30.3888 38 21C38 11.6112 30.3888 4 21 4C11.6112 4 4 11.6112 4 21C4 30.3888 11.6112 38 21 38Z"/><path stroke="#fff" stroke-linecap="round" d="M21 15L21 27"/><path stroke="#fff" stroke-linecap="round" d="M15.0156 21.0156L27 21"/><path stroke="#000" stroke-linecap="round" d="M33.2216 33.2217L41.7069 41.707"/></g>`),
		g.Group(children),
	)
}