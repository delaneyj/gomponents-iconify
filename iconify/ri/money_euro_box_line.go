package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyEuroBoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.005 3.003h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1Zm1 2v14h16v-14h-16Zm6.05 6h4.95v2h-4.95a2.5 2.5 0 0 0 4.064 1.41l1.7 1.133a4.5 4.5 0 0 1-7.787-2.543H7.005v-2h1.027A4.5 4.5 0 0 1 15.82 8.46l-1.701 1.134a2.5 2.5 0 0 0-4.064 1.41Z"/>`),
		g.Group(children),
	)
}