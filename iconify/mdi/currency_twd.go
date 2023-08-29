package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyTwd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11h18v2h-6v6h6v2h-6a2 2 0 0 1-2-2v-6h-2.65l-4.62 8L4 20l4.04-7H3v-2m2-8h14v2H5V3Z"/>`),
		g.Group(children),
	)
}