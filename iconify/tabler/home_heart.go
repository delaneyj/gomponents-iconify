package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 12l-9-9l-9 9h2v7a2 2 0 0 0 2 2h6"/><path d="M9 21v-6a2 2 0 0 1 2-2h2c.39 0 .754.112 1.061.304M19 21.5l2.518-2.58a1.74 1.74 0 0 0 0-2.413a1.627 1.627 0 0 0-2.346 0l-.168.172l-.168-.172a1.627 1.627 0 0 0-2.346 0a1.74 1.74 0 0 0 0 2.412l2.51 2.59z"/></g>`),
		g.Group(children),
	)
}