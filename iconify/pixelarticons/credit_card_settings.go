package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4H2v16h18v-2H4v-6h16V8H4V6h16V4zm0 0h2v16h-2V4zm-7 18h-2v2h2v-2zm2 0h2v2h-2v-2zm-6 0H7v2h2v-2z"/>`),
		g.Group(children),
	)
}