package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardRefundOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2m0 10H6v-4h12v4m0-6H6v-2h12v2M17 5v5h-1.5V6.5H9.88l2.42 2.43L11.24 10L7 5.75l4.24-4.25l1.06 1.07L9.88 5H17Z"/>`),
		g.Group(children),
	)
}