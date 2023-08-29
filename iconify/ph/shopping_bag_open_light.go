package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagOpenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42H40a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14ZM40 54h176a2 2 0 0 1 2 2v18H38V56a2 2 0 0 1 2-2Zm176 148H40a2 2 0 0 1-2-2V86h180v114a2 2 0 0 1-2 2Zm-42-90a46 46 0 0 1-92 0a6 6 0 0 1 12 0a34 34 0 0 0 68 0a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}