package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M14 38h20v6H14zm4 0l2-29l4-5l4 5l2 29H18Z"/><path stroke-linecap="round" d="M4 44h40"/></g>`),
		g.Group(children),
	)
}