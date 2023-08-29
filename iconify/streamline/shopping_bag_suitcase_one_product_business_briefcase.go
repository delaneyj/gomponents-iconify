package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagSuitcaseOneProductBusinessBriefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10" x=".5" y="3.5" rx="1"/><path d="M5 .5h4a1 1 0 0 1 1 1v2h0h-6h0v-2a1 1 0 0 1 1-1ZM3.5 7h7m-7 3h7"/></g>`),
		g.Group(children),
	)
}