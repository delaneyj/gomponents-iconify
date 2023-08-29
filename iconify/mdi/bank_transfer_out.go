package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTransferOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15v-3h3v-2l4 3.5l-4 3.5v-2h-3m-1-6.3V10H2V8.7L8 5l6 3.7M2 17h12v2H2v-2m5-6h2v5H7v-5m-4 0h2v5H3v-5m8 0h2v5h-2v-5Z"/>`),
		g.Group(children),
	)
}