package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1Zm3 10v2h2v-2H7Zm0 4v2h2v-2H7Zm4-4v2h2v-2h-2Zm0 4v2h2v-2h-2Zm4-4v6h2v-6h-2ZM7 6v4h10V6H7Z"/>`),
		g.Group(children),
	)
}