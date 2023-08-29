package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirBike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40M28 30h6.19c2.27 0 6.81 1.344 6.81 6.72V44m-6-14l5-11l-6-13m-5 2l10-4"/><circle cx="20" cy="30" r="8"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 30h8m-7-8l-7-9m-4 0h8"/><path d="M20 38a8 8 0 1 0 0-16"/></g>`),
		g.Group(children),
	)
}