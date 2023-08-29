package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EngineeringBrand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="38" height="20" x="5" y="6" rx="2"/><path stroke-linecap="round" d="M12 26v16m24-9H12m4 9H8m32 0h-8m4-16v16"/></g>`),
		g.Group(children),
	)
}