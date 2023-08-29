package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SentToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M30 18h11a1 1 0 0 1 1 1v22a1 1 0 0 1-1 1H19a1 1 0 0 1-1-1V30M9.97 6H6v4.034M9.997 30H6v-3.988"/><path stroke-linejoin="round" d="M26 30h3.997v-3.988M26.002 6H30v3.998"/><path d="M16.028 6h3.98"/><path stroke-linejoin="round" d="M6 16v4.015M30 16v4.015M15.992 30H20"/></g>`),
		g.Group(children),
	)
}