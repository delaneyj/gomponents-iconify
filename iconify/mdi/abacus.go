package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abacus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h2v6H5V5m5 0H8v6h2V5M5 19h2v-6H5v6m5-6H8v6h2v-2h5v-2h-5v-2m-8 8h2V3H2v18M20 3v4h-7V5h-2v6h2V9h7v6h-2v-2h-2v6h2v-2h2v4h2V3h-2Z"/>`),
		g.Group(children),
	)
}