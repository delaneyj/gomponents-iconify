package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBatteryTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17v5h10v-5H7z"/><path fill="currentColor" fill-opacity=".3" d="M17 4h-3V2h-4v2H7v13h10V4z"/>`),
		g.Group(children),
	)
}