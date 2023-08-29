package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H14"/><path d="M20 8v14H6.5a2.5 2.5 0 0 1 0-5H20"/><circle cx="14" cy="8" r="2"/><path d="m20 2l-4.5 4.5M19 3l1 1"/></g>`),
		g.Group(children),
	)
}