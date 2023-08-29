package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trolley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-3-3l3 2m3-1l8-12m-3 5l2 1M9.592 4.695l3.306 2.104a1.3 1.3 0 0 1 .396 1.8L10.2 13.41a1.3 1.3 0 0 1-1.792.394L5.102 11.7a1.3 1.3 0 0 1-.396-1.8L7.8 5.09a1.3 1.3 0 0 1 1.792-.394z"/>`),
		g.Group(children),
	)
}