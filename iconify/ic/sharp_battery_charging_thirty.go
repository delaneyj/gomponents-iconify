package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBatteryChargingThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M17 4h-3V2h-4v2H7v10.5h2L13 7v5.5h2l-1.07 2H17V4z"/><path fill="currentColor" d="M11 20v-5.5H7V22h10v-7.5h-3.07L11 20z"/>`),
		g.Group(children),
	)
}