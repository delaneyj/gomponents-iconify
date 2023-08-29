package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path stroke-linecap="round" d="m15.5 9l.172.172c1.333 1.333 2 2 2 2.828c0 .828-.667 1.495-2 2.828L15.5 15m-2.206-7.83L12 12l-1.294 4.83M8.5 9l-.171.172c-1.334 1.333-2 2-2 2.828c0 .828.666 1.495 2 2.828L8.5 15"/></g>`),
		g.Group(children),
	)
}