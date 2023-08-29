package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSignal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 22v-2m3 2v-4m3 4v-6m-2-3.506V12h2l-9-9l-9 9h2v7a2 2 0 0 0 2 2h4"/><path d="M9 21v-6a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v.5"/></g>`),
		g.Group(children),
	)
}