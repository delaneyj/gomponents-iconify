package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44H40a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm4 156a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4V56a4 4 0 0 1 4-4h176a4 4 0 0 1 4 4ZM172 88a44 44 0 0 1-88 0a4 4 0 0 1 8 0a36 36 0 0 0 72 0a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}