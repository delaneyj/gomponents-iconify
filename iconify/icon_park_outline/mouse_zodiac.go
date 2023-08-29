package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 35H4c0-8 5-18 10-18l2 9m12 9c0-5 2-8.01 9-8"/><path d="M44 28.537C45 33.511 42 35 40 34s-1.5-6-3-10c-3.14-8.375-15-10-22-3"/></g>`),
		g.Group(children),
	)
}