package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Communication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M33 38H22v-8h14v-8h8v16h-5l-3 3l-3-3Z"/><path stroke-linejoin="round" d="M4 6h32v24H17l-4 4l-4-4H4V6Z"/><path d="M19 18h1m6 0h1m-15 0h1"/></g>`),
		g.Group(children),
	)
}