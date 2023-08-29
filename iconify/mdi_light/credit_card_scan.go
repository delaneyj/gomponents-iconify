package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6h11a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2h15a2 2 0 0 0-2-2H6m-2 9a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2v-4H4v4m2-1h3v1H6v-1m5 0h2v1h-2v-1m-7-5v1h15v-1H4M3 3h2v1H3a2 2 0 0 0-2 2v2H0V6a3 3 0 0 1 3-3M1 19a2 2 0 0 0 2 2h2v1H3a3 3 0 0 1-3-3v-2h1v2M20 3a3 3 0 0 1 3 3v2h-1V6a2 2 0 0 0-2-2h-2V3h2m3 16a3 3 0 0 1-3 3h-2v-1h2a2 2 0 0 0 2-2v-2h1v2Z"/>`),
		g.Group(children),
	)
}