package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAreaStepper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 10V6H10v6H4V2H2v26a2.002 2.002 0 0 0 2 2h26V10Zm-10 4V8h8v4h8v10h-6v-6H12v6H4v-8ZM4 28v-4h10v-6h6v6h8v4Z"/>`),
		g.Group(children),
	)
}