package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h18v2H4v2h16v4H4v6h10v2H2V4zm20 0h-2v8h2V4zm-4 10h2v2h2v2h-2v2h-2v-2h-2v-2h2v-2z"/>`),
		g.Group(children),
	)
}