package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayBasketball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 4a1 1 0 1 0 2 0a1 1 0 0 0-2 0M5 21l3-3l.75-1.5M14 21v-4l-4-3l.5-6"/><path d="m5 12l1-3l4.5-1l3.5 3l4 1"/><path fill="currentColor" d="M18.5 16a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}