package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.44 5.14H4.56a2.5 2.5 0 0 0-2.5 2.5v8.72a2.5 2.5 0 0 0 2.5 2.5h14.88a2.5 2.5 0 0 0 2.5-2.5V7.64a2.5 2.5 0 0 0-2.5-2.5ZM3.06 7.64a1.5 1.5 0 0 1 1.5-1.5h14.88a1.5 1.5 0 0 1 1.5 1.5v.5H3.06Zm17.88 8.72a1.5 1.5 0 0 1-1.5 1.5H4.56a1.5 1.5 0 0 1-1.5-1.5V9.64h17.88Z"/><path fill="currentColor" d="M8.063 14.247h-3a.5.5 0 1 1 0-1h3a.5.5 0 1 1 0 1Zm10.871.003h-6.5a.5.5 0 1 1 0-1h6.5a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}