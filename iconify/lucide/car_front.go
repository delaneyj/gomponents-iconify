package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 8l-2 2l-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10L3 8m4 6h.01M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2m14-2v2"/></g>`),
		g.Group(children),
	)
}