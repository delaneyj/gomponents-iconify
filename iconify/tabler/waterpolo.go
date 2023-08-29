package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waterpolo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 9a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/><path d="m5 8l3 4l4.5 1l7.5-1M3 18.75A2.4 2.4 0 0 0 4 19a2.4 2.4 0 0 0 2-1a2.4 2.4 0 0 1 2-1a2.4 2.4 0 0 1 2 1a2.4 2.4 0 0 0 2 1a2.4 2.4 0 0 0 2-1a2.4 2.4 0 0 1 2-1a2.4 2.4 0 0 1 2 1a2.4 2.4 0 0 0 2 1a2.4 2.4 0 0 0 1-.25M12 16l.5-3"/><path fill="currentColor" d="M6.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}