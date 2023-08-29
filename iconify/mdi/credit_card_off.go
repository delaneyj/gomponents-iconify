package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.2 8l-4-4H20a2 2 0 0 1 2 2v12a1.91 1.91 0 0 1-.12.68L14.2 11H20V8m.84 14.73L18.11 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 .65-1.46L1.11 3l1.28-1.27l19.72 19.73M9.11 11l-3-3H4v3Z"/>`),
		g.Group(children),
	)
}