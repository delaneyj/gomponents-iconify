package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingBroadcastTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/><path d="M16.616 13.924a5 5 0 1 0-9.23 0"/><path d="M20.307 15.469a9 9 0 1 0-16.615 0"/><path d="m9 21l3-9l3 9m-5-2h4"/></g>`),
		g.Group(children),
	)
}