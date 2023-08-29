package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagAltF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6V3a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v3h2v11a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3V6h2zm2 0h6V3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v3z"/>`),
		g.Group(children),
	)
}