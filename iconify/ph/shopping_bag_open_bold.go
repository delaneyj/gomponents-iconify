package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagOpenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v144a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 24v16H44V60ZM44 196v-96h168v96Zm136-72a52 52 0 0 1-104 0a12 12 0 0 1 24 0a28 28 0 0 0 56 0a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}