package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayVolleyball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 4a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/><path fill="currentColor" d="M20.5 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/><path d="m2 16l5 1l.5-2.5m4 6.5l2.5-5.5L8.5 12L12 8l3 4l4 2"/></g>`),
		g.Group(children),
	)
}