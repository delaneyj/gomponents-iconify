package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M14 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path stroke-linecap="round" d="M18 10s-3.537 1.5-6 1.5S6 10 6 10m6 2v1.452m0 0a3 3 0 0 0 .476 1.623L15 19m-3-5.548a3 3 0 0 1-.476 1.623L9 19"/></g>`),
		g.Group(children),
	)
}