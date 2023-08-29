package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyFlorin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8m-9 7c1.213 0 2.31-.723 2.788-1.838l4.424-10.324A3.033 3.033 0 0 1 17 5"/>`),
		g.Group(children),
	)
}