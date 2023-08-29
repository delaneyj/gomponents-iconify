package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeautyInstrument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 34h14m-15-7h16"/><rect width="30" height="16" x="9" y="4" rx="4"/><path d="m14 20l5 24h10l5-24m-17-8h13"/></g>`),
		g.Group(children),
	)
}