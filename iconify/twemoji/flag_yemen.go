package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagYemen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#141414" d="M0 27.063C0 29.272 1.791 31 4 31h28c2.209 0 4-1.728 4-3.937V22H0v5.063z"/><path fill="#EEE" d="M0 14h36v8H0z"/><path fill="#CE1126" d="M32 5H4C1.791 5 0 6.854 0 9.063V14h36V9.063C36 6.854 34.209 5 32 5z"/>`),
		g.Group(children),
	)
}