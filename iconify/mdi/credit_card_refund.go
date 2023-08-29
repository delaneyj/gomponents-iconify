package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardRefund(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2m0 6H6v-3h12m-1-9v5h-1.5V6.5H9.88l2.42 2.43L11.24 10L7 5.75l4.24-4.25l1.06 1.07L9.88 5Z"/>`),
		g.Group(children),
	)
}