package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbMicroOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 18a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V18Z"/><path d="M11 15h26v8H11v-8Zm10 8v-2m6 2v-2m5 2v-2m-16 2v-2m-8-6h32"/></g>`),
		g.Group(children),
	)
}