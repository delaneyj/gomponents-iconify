package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 8l20 13v23H4V21L24 8Z" clip-rule="evenodd"/><path d="M20 44V23l-8 5v16m16 0V23l8 5v16m5 0H8"/></g>`),
		g.Group(children),
	)
}