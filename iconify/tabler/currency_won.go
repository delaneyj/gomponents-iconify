package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyWon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 6l3.245 11.358a.85.85 0 0 0 1.624.035L12 8l3.131 9.393a.85.85 0 0 0 1.624-.035L20 6m1 4H3m18 4H3"/>`),
		g.Group(children),
	)
}