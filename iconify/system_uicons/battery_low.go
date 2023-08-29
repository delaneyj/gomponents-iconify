package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.497h10a2 2 0 0 1 2 2V11.5a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2V8.497a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5 8a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 8.5v3"/></g>`),
		g.Group(children),
	)
}