package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingStoreSignageThreeStreetSandwichShopsShopStoresBoardSignStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 2.5h-7l-3 11h7l3-11zm3 11l-3-11m-2.18 8h4.36"/>`),
		g.Group(children),
	)
}