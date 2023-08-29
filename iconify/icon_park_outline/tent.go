package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 12L4 36h12"/><path d="M38 12H10l6 24h28l-6-24Zm-26 6h27m-29-6l3 12m25-12l3 12"/></g>`),
		g.Group(children),
	)
}