package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 19a3 3 0 1 0 6 0a3 3 0 1 0-6 0m1 2l4-4m-2-5h2l-9-9l-9 9h2v7a2 2 0 0 0 2 2h5.5"/><path d="M9 21v-6a2 2 0 0 1 2-2h2c.58 0 1.103.247 1.468.642"/></g>`),
		g.Group(children),
	)
}