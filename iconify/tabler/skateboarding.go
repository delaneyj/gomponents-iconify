package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skateboarding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 4a1 1 0 1 0 2 0a1 1 0 0 0-2 0M5.5 15H9l.75-1.5M14 19v-5l-2.5-3L14 7"/><path d="m8 8l3-1h4l1 3h3"/><path fill="currentColor" d="M17.5 21a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/><path d="M3 18c0 .552.895 1 2 1h14c1.105 0 2-.448 2-1"/><path fill="currentColor" d="M6.5 21a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}