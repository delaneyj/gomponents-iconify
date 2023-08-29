package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTransferIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 15v-3h3v-2l4 3.5L5 17v-2H2m20-6.3V10H10V8.7L16 5l6 3.7M10 17h12v2H10v-2m5-6h2v5h-2v-5m-4 0h2v5h-2v-5m8 0h2v5h-2v-5Z"/>`),
		g.Group(children),
	)
}