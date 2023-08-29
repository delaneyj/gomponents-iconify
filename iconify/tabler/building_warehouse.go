package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingWarehouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 21V8l9-4l9 4v13"/><path d="M13 13h4v8H7v-6h6"/><path d="M13 21v-9a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3"/></g>`),
		g.Group(children),
	)
}