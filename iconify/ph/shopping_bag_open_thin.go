package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagOpenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 44H40a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12ZM40 52h176a4 4 0 0 1 4 4v20H36V56a4 4 0 0 1 4-4Zm176 152H40a4 4 0 0 1-4-4V84h184v116a4 4 0 0 1-4 4Zm-44-92a44 44 0 0 1-88 0a4 4 0 0 1 8 0a36 36 0 0 0 72 0a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}