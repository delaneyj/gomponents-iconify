package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBatteryChargingFifty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.47 13.5L11 20v-5.5H9l.53-1H7V22h10v-8.5h-2.53z"/><path fill="currentColor" fill-opacity=".3" d="M17 4h-3V2h-4v2H7v9.5h2.53L13 7v5.5h2l-.53 1H17V4z"/>`),
		g.Group(children),
	)
}