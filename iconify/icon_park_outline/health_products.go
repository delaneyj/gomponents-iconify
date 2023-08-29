package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthProducts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M33 12H15l-5 5.843v20.249L15 44h18l5-5.908v-20.25L33 12Zm-14 8h10m4-8V7a3 3 0 0 0-3-3H18a3 3 0 0 0-3 3v5"/><circle cx="24" cy="32" r="5"/></g>`),
		g.Group(children),
	)
}