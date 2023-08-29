package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M4 39h40M4 27h11.5m17 0H44"/><path d="M15.514 27a11.021 11.021 0 0 1-.735-1A10.949 10.949 0 0 1 13 20c0-6.075 4.925-11 11-11s11 4.925 11 11a10.949 10.949 0 0 1-2.514 7"/></g>`),
		g.Group(children),
	)
}