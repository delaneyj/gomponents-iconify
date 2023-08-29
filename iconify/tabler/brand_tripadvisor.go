package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTripadvisor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 13.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m11 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0"/><path d="M17.5 9a4.5 4.5 0 1 0 3.5 1.671L22 9h-4.5zm-11 0A4.5 4.5 0 1 1 3 10.671L2 9h4.5z"/><path d="m10.5 15.5l1.5 2l1.5-2M9 6.75c2-.667 4-.667 6 0"/></g>`),
		g.Group(children),
	)
}