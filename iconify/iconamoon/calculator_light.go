package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M5 9h14"/><path stroke-linejoin="round" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V3Z"/><path stroke-linecap="round" d="M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`),
		g.Group(children),
	)
}