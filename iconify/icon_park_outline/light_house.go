package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 44h36M17 20h14"/><path d="M19 20h10l3 24H16l3-24Z"/><path stroke-linecap="round" d="m19 9l2 11h6l2-11"/><path stroke-linecap="round" d="m32 12l-3-3l-5-5l-5 5l-3 3m21-5l3-3M11 7L8 4m29 15l3 3m-29-3l-3 3m30-9h4m-32 0H6m12 15h12m-13 8h14"/><path d="m29 20l3 24M19 20l-3 24"/></g>`),
		g.Group(children),
	)
}