package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingStoreSignageTwoSignStoreShopShopsSignageStoresPole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="7" x=".5" y=".5" rx="2"/><path d="M7 13.5v-6"/></g>`),
		g.Group(children),
	)
}