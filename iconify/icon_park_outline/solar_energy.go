package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarEnergy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="40" height="24" x="4" y="8" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 32V8M18 32V8m24 12H6m18 21v-9m7 9H17"/></g>`),
		g.Group(children),
	)
}