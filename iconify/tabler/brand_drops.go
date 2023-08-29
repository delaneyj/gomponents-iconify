package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDrops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.637 7.416a7.907 7.907 0 0 1 1.76 8.666A8 8 0 0 1 12 21a8 8 0 0 1-7.396-4.918a7.907 7.907 0 0 1 1.759-8.666L12 2l5.637 5.416z"/><path d="M14.466 10.923a3.595 3.595 0 0 1 .77 3.877A3.5 3.5 0 0 1 12 17a3.5 3.5 0 0 1-3.236-2.2a3.595 3.595 0 0 1 .77-3.877L12 8.5l2.466 2.423z"/></g>`),
		g.Group(children),
	)
}