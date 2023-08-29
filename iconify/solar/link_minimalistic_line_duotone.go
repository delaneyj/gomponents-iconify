package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M9 12h6"/><path d="M9 18H8A6 6 0 0 1 8 6h1m6 0h1a6 6 0 0 1 0 12h-1" opacity=".5"/></g>`),
		g.Group(children),
	)
}