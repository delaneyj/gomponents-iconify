package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDeliveroo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 11l1-9l5 .5L20 16l-3 6l-12.5-2.5l-1.5-6l7-1.5l-1.5-7.5l4.5-1z"/><circle cx="15.5" cy="15.5" r="1" fill="currentColor"/><circle cx="11.5" cy="14.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}