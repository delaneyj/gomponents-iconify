package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarEnergyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M10 30h6a6 6 0 0 1 0 12h-6V30Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 28v16M4 32h6m-6 6h6m12-2h8a2 2 0 0 0 2-2V23m0-19v7m-9.121-3.121l4.242 4.242m-4.242 13l4.242-4.242m14-13L36.88 12.12m4.241 13.001L36.88 20.88"/><circle cx="32" cy="17" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 17h6m12 0h6"/></g>`),
		g.Group(children),
	)
}