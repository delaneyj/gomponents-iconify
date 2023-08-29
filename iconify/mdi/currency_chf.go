package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyChf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M7 3h11v2H9v6h8v2H9v3h2v2H9v3H7v-3H5v-2h2V3z" fill="currentColor"/>`),
		g.Group(children),
	)
}