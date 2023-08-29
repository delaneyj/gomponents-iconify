package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h18v1h1v12h-1v1H2v-1H1V5h1V4m1 2v2h16V6H3m0 10h16v-5H3v5Z"/>`),
		g.Group(children),
	)
}