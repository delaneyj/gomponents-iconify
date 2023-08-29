package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkiingNordic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-miterlimit="2" stroke-width="4"><path d="M24 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 24h-6l-6-5l-3 10l6 6l2 9M17 34l-2 5l-6 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 44h34l5-7M34 24v20"/></g>`),
		g.Group(children),
	)
}