package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCurrencyYuan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.28 12H18v2h-5v7h-2v-7H6v-2h4.72L5 3h2.37L12 10.29L16.63 3H19z"/>`),
		g.Group(children),
	)
}