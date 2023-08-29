package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Karate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 9l4.5 1l3 2.5M13 21v-8l3-5.5"/><path d="m8 4.5l4 2l4 1l4 3.5l-2 3.5"/></g>`),
		g.Group(children),
	)
}