package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDisney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.22 5.838C1.913 5.688 2 5.26 2 5.044C2 4.828 2.424 4 6.34 4C11.034 4 21 7.645 21 14.042s-8.71 4.931-10.435 4.52C8.841 18.15 5 16.306 5 14.388C5 12.993 8.08 12 11.715 12C15.349 12 17 13.041 17 14c0 .5-.074 1.229-1 1.5"/><path d="M10.02 8a505.153 505.153 0 0 0 0 13"/></g>`),
		g.Group(children),
	)
}