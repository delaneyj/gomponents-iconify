package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EngineeringVehicle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M32 6h6M10 36H6v-8h26v8H18m14 0V12h6.5L44 24v12h-3"/><path stroke-linejoin="round" d="M4 8h22v14H7L4 8Z"/><circle cx="37" cy="38" r="4"/><circle cx="14" cy="38" r="4"/></g>`),
		g.Group(children),
	)
}