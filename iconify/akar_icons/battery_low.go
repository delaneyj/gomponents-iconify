package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="2" y="6" rx="2"/><path d="M7 10v4m13-4h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H20v-4Z"/></g>`),
		g.Group(children),
	)
}