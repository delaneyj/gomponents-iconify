package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5m4-1v5l4 2"/></g>`),
		g.Group(children),
	)
}