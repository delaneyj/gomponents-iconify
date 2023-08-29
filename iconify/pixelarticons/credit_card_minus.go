package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4H2v16h12v-2H4v-6h16V8H4V6h16V4zm0 0h2v8h-2V4zm2 12h-6v2h6v-2z"/>`),
		g.Group(children),
	)
}