package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVscode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 3v18l4-2.5v-13zM9.165 13.903L5 17.5l-2-1L7.333 12m1.735-1.802L16 3v5l-4.795 4.141"/><path d="M16 16.5L5 6.5l-2 1L16 21"/></g>`),
		g.Group(children),
	)
}