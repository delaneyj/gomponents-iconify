package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBatteryChargingNinety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M17 4h-3V2h-4v2H7v4h5.47L13 7v1h4V4z"/><path fill="currentColor" d="M13 12.5h2L11 20v-5.5H9L12.47 8H7v14h10V8h-4v4.5z"/>`),
		g.Group(children),
	)
}