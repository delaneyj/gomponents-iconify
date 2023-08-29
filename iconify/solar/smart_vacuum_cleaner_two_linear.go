package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartVacuumCleanerTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/><path d="M15 9a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path stroke-linecap="round" d="M12 18v-2M2 22l2.5-2.5M22 22l-2.5-2.5"/></g>`),
		g.Group(children),
	)
}