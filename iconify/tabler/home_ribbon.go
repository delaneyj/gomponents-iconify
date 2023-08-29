package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRibbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 15h5v7l-2.5-1.5L16 22z"/><path d="m20 11l-8-8l-9 9h2v7a2 2 0 0 0 2 2h5"/><path d="M9 21v-6a2 2 0 0 1 2-2h1.5"/></g>`),
		g.Group(children),
	)
}