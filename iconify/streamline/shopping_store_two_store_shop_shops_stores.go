package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingStoreTwoStoreShopShopsStores(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 8.5V13a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8.5M13 6H1a.5.5 0 0 1-.5-.5v-2l1.22-2.45a1 1 0 0 1 .9-.55h8.76a1 1 0 0 1 .9.55L13.5 3.5v2a.5.5 0 0 1-.5.5ZM8 8.5v5M1.5 10H8M.5 3.5h13"/>`),
		g.Group(children),
	)
}