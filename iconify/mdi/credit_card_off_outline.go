package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m.93 4.2l1.28-1.27L20 20.72L18.73 22l-2-2H4a2 2 0 0 1-2-2V6c0-.22.04-.43.11-.62L.93 4.2M20 8V6H7.82l-2-2H20a2 2 0 0 1 2 2v12c0 .6-.26 1.13-.68 1.5l-1.5-1.5H20v-6h-6.18l-4-4H20M4 8h.73L4 7.27V8m0 4v6h10.73l-6-6H4Z"/>`),
		g.Group(children),
	)
}