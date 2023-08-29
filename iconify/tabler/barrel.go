package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barrel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.278 4h9.444a2 2 0 0 1 1.841 1.22C19.521 7.48 20 9.74 20 12c0 2.26-.479 4.52-1.437 6.78A2 2 0 0 1 16.722 20H7.278a2 2 0 0 1-1.841-1.22C4.479 16.52 4 14.26 4 12c0-2.26.479-4.52 1.437-6.78A2 2 0 0 1 7.278 4z"/><path d="M14 4c.667 2.667 1 5.333 1 8s-.333 5.333-1 8M10 4c-.667 2.667-1 5.333-1 8s.333 5.333 1 8m-5.5-4h15m0-8h-15"/></g>`),
		g.Group(children),
	)
}