package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 21v-6a2 2 0 0 1 2-2h2c.641 0 1.212.302 1.578.771"/><path d="M20.136 11.136L12 3l-9 9h2v7a2 2 0 0 0 2 2h6.344M19 22v-6m3 3l-3-3l-3 3"/></g>`),
		g.Group(children),
	)
}