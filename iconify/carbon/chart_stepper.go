package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStepper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 22v-8h8V6h8V4H20v8h-8v8H4V2H2v26a2.002 2.002 0 0 0 2 2h26v-2H4v-6Z"/>`),
		g.Group(children),
	)
}