package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingCircus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 11h16m-8-4.5c0 1-5 4.5-8 4.5m8-4.5c0 1 5 4.5 8 4.5M6 11c-.333 5.333-1 8.667-2 10h4c1 0 4-4 4-9v-1m6 0c.333 5.333 1 8.667 2 10h-4c-1 0-4-4-4-9v-1"/><path d="M12 7V3l2 1h-2"/></g>`),
		g.Group(children),
	)
}