package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCatergoriesChairDesignLoungeFurnitureChairInteriorDecorateArmchairDecoration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 7.5h1a1 1 0 0 1 1 1v5h0h-3h0v-5a1 1 0 0 1 1-1Zm10 0h1a1 1 0 0 1 1 1v5h0h-3h0v-5a1 1 0 0 1 1-1Zm-8 6h7m-8-6v-2a3 3 0 0 1 3-3h3a3 3 0 0 1 3 3v2m-8 2h7"/>`),
		g.Group(children),
	)
}