package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H10"/><path d="M20 15v7H6.5a2.5 2.5 0 0 1 0-5H20"/><rect width="8" height="5" x="12" y="6" rx="1"/><path d="M18 6V4a2 2 0 1 0-4 0v2"/></g>`),
		g.Group(children),
	)
}