package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabaledWeb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M25 40H7a3 3 0 0 1-3-3V11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v12.059"/><circle cx="38" cy="35" r="6" stroke="currentColor" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m34 31l8 8"/><path stroke="currentColor" stroke-width="4" d="M4 11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v9H4v-9Z"/><circle r="2" fill="currentColor" transform="matrix(0 -1 -1 0 10 14)"/><circle r="2" fill="currentColor" transform="matrix(0 -1 -1 0 16 14)"/></g>`),
		g.Group(children),
	)
}