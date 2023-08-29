package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42.5 16.387a20.018 20.018 0 0 0-4-6.162A19.943 19.943 0 0 0 24 4C12.954 4 4 12.954 4 24h10l5 8l9-18l7 10h9c0 11.046-8.954 20-20 20c-5.45 0-10.393-2.18-14-5.717A20.04 20.04 0 0 1 5.664 32"/>`),
		g.Group(children),
	)
}