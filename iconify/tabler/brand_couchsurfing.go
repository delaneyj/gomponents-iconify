package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCouchsurfing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.1 13c3.267 0 5.9-.167 7.9-.5c3-.5 4-2 4-3.5a3 3 0 1 0-6 0c0 1.554 1.807 3 3 4s2 2.5 2 3.5a1.5 1.5 0 1 1-3 0c0-2 4-3.5 7-3.5h2.9"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/></g>`),
		g.Group(children),
	)
}