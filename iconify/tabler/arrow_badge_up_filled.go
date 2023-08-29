package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBadgeUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m11.375 6.22l-5 4A1 1 0 0 0 6 11v6l.006.112a1 1 0 0 0 1.619.669L12 14.28l4.375 3.5A1 1 0 0 0 18 17v-6a1 1 0 0 0-.375-.78l-5-4a1 1 0 0 0-1.25 0z"/></g>`),
		g.Group(children),
	)
}