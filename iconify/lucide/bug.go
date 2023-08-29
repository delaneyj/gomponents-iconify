package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="14" x="8" y="6" rx="4"/><path d="m19 7l-3 2M5 7l3 2m11 10l-3-2M5 19l3-2m12-4h-4M4 13h4m2-9l1 2m3-2l-1 2"/></g>`),
		g.Group(children),
	)
}