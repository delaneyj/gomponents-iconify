package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M3 12c5.4-8 12.6-8 18 0c-5.4 8-12.6 8-18 0z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/></g>`),
		g.Group(children),
	)
}