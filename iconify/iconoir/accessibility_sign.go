package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilitySign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.5 12.5l7-.5l-1.5 6.5m-5.5-6l4.5-5L12.5 5L10 7.5m8.5-1a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path d="M5.5 12.5a5 5 0 0 1 7.584 6M3.729 15A5 5 0 0 0 11 20.831"/></g>`),
		g.Group(children),
	)
}