package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 5.835A10.485 10.485 0 0 0 12 1.5c-5.427 0-9.89 4.115-10.443 9.396l-.104.994l1.99.209l.103-.995A8.501 8.501 0 0 1 19.212 7.5H17.5v2h5v-7h-2v3.335ZM11 6v1a3 3 0 0 0 0 6h2a1 1 0 1 1 0 2H8.5v2H11v1h2v-1a3 3 0 1 0 0-6h-2a1 1 0 0 1 0-2h4.5V7H13V6h-2Zm9.557 5.901l-.104.995A8.501 8.501 0 0 1 4.787 16.5H6.5v-2h-5v7h2v-3.335A10.485 10.485 0 0 0 12 22.5c5.426 0 9.89-4.115 10.442-9.396l.104-.994l-1.989-.209Z"/>`),
		g.Group(children),
	)
}