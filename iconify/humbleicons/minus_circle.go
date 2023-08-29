package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M16 12H8"/><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/></g>`),
		g.Group(children),
	)
}