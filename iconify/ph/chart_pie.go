package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm71.87 53.27L136 114.14V40.37a88 88 0 0 1 63.87 36.9ZM120 40.37v83l-71.89 41.5A88 88 0 0 1 120 40.37ZM128 216a88 88 0 0 1-71.87-37.27l151.76-87.61A88 88 0 0 1 128 216Z"/>`),
		g.Group(children),
	)
}