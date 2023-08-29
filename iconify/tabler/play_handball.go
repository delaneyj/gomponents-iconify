package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayHandball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13 21l3.5-2l-4.5-4l2-4.5"/><path d="m7 6l2 4l5 .5l4 2.5l2.5 3M4 20l5-1l1.5-2M15 7a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/><path fill="currentColor" d="M9.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}