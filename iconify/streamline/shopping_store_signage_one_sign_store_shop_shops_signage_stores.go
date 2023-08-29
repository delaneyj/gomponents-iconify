package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingStoreSignageOneSignStoreShopShopsSignageStores(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="2"/><path d="m4 5.5l3-5l3 5"/></g>`),
		g.Group(children),
	)
}