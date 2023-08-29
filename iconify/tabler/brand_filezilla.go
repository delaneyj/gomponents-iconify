package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandFilezilla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 15.824a4.062 4.062 0 0 1-2.25.033c-.738-.201-2.018-.08-2.75.143l4.583-5H9"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="m8 15l2-8h5"/></g>`),
		g.Group(children),
	)
}