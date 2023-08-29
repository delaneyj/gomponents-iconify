package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M41 5.5H7v28h34v-28Z"/><path stroke-linecap="round" d="m16 41.5l8-8l8 8M13.924 24.663l5.642-5.508l4.442 4.345l9.959-9.98M4 5.5h40"/></g>`),
		g.Group(children),
	)
}