package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 24c0-11.046-8.954-20-20-20S4 12.954 4 24h40Z"/><path stroke-linecap="round" d="M24 24v14.554c0 3.014 2.486 5.457 5.5 5.457s5.5-2.443 5.5-5.457"/></g>`),
		g.Group(children),
	)
}