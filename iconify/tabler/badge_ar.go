package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeAr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M7 15v-4.5a1.5 1.5 0 0 1 3 0V15m-3-2h3m4-1h1.5a1.5 1.5 0 0 0 0-3H14v6m3 0l-2-3"/></g>`),
		g.Group(children),
	)
}