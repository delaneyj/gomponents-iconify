package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeEco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20 11l-8-8l-9 9h2v7a2 2 0 0 0 2 2h5"/><path d="M9 21v-6a2 2 0 0 1 2-2h2c.325 0 .631.077.902.215M16 22s0-2 3-4"/><path d="M19 21a3 3 0 0 1 0-6h3v3a3 3 0 0 1-3 3z"/></g>`),
		g.Group(children),
	)
}