package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathSymbols(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18m-9-9v18m4.5-16.5l3 3m0-3l-3 3M6 4v4M4 6h4m10 10h.01M18 20h.01M4 18h4"/>`),
		g.Group(children),
	)
}