package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircuitVoltmeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12a7 7 0 1 0 14 0a7 7 0 1 0-14 0m0 0H2m17 0h3"/><path d="m10 10l2 4l2-4"/></g>`),
		g.Group(children),
	)
}