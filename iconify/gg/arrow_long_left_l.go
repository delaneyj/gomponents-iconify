package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongLeftL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.208 7.757L.97 12.003l4.246 4.24l1.413-1.416l-1.819-1.815l16.214.018l-.004 2l2 .004l.012-6l-2-.004l-.006 2.989l.001-.99l-16.24-.018l1.838-1.84l-1.416-1.414Z"/>`),
		g.Group(children),
	)
}