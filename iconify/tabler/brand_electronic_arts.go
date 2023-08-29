package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandElectronicArts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="m17.5 15l-3-6l-3 6h-5L8 12m9 2h-2m-8.5-2H10M8 9h3"/></g>`),
		g.Group(children),
	)
}