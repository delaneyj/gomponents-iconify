package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBatteryChargingEighty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M17 4h-3V2h-4v2H7v5h4.93L13 7v2h4V4z"/><path fill="currentColor" d="M13 12.5h2L11 20v-5.5H9L11.93 9H7v13h10V9h-4v3.5z"/>`),
		g.Group(children),
	)
}