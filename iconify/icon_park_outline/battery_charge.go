package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 14H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h7m18 0h7a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2h-7m-8.998 0L17 24.001h10.004L22 34"/><path fill="currentColor" d="M42 20h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2v-8Z"/></g>`),
		g.Group(children),
	)
}