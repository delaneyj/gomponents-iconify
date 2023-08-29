package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandOnlyfans(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.5 6a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13z"/><path d="M8.5 15a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5zm5.5 1c2.5 0 6.42-1.467 7-4h-6c3-1 6.44-3.533 7-6h-4c-3.03 0-3.764-.196-5 1.5"/></g>`),
		g.Group(children),
	)
}