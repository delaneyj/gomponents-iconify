package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10.5v3M8.5 17H6a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h.5m9 10h.5a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-2.5M10 7l-1 5.5l4-1l-1 5.5"/>`),
		g.Group(children),
	)
}