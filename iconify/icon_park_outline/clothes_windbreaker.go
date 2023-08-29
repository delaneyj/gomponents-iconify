package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesWindbreaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 25v12m-26 0v7h26v-7m-26 0H4V18l6.125-5L17 18l7-8l7 8l6.875-5L44 18v19h-7m-26 0V25"/><path d="m17 18l7-8l7 8l11-8l-2.5-6H9l-3 6l11 8Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 10h5m5 0h-5m0 0l-7 8l-11-8l3-6h30.5l2.5 6l-11 8l-7-8Z"/></g>`),
		g.Group(children),
	)
}