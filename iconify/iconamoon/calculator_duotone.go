package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 9h14v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V9Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 9h14"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V3Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`),
		g.Group(children),
	)
}