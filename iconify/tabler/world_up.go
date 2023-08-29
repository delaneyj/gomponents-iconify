package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.985 12.52a9 9 0 1 0-8.451 8.463M3.6 9h16.8M3.6 15h10.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18a16.996 16.996 0 0 1 2.391 11.512M19 22v-6m3 3l-3-3l-3 3"/></g>`),
		g.Group(children),
	)
}