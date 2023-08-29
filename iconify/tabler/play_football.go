package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayFootball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4a1 1 0 1 0 2 0a1 1 0 0 0-2 0M3 17l5 1l.75-1.5M14 21v-4l-4-3l1-6"/><path d="M6 12V9l5-1l3 3l3 1"/><path fill="currentColor" d="M19.5 20a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}