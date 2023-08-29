package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8V5a5 5 0 1 1 10 0v3h2a3 3 0 0 1 3 3v3a6 6 0 0 1-6 6H6a6 6 0 0 1-6-6v-3a3 3 0 0 1 3-3h2zm2 0h6V5a3 3 0 0 0-6 0v3z"/>`),
		g.Group(children),
	)
}