package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFailure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="36" height="20" x="14" y="44" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2" transform="rotate(-90 14 44)"/><path fill="currentColor" d="M20 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2h-8Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 30v-3c2.21 0 4-2.015 4-4.5S26.21 18 24 18s-4 2.015-4 4.5"/><path fill="currentColor" stroke="currentColor" d="M26 35.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g>`),
		g.Group(children),
	)
}