package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M38 20H18v8h20v-8ZM32 6H18v8h14V6Zm12 28H18v8h26v-8Z"/><path stroke-linecap="round" d="M17 10H5m12 14H5m12 14H5m0 6V4"/></g>`),
		g.Group(children),
	)
}