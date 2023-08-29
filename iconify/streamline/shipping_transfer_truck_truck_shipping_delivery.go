package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShippingTransferTruckTruckShippingDelivery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 5.5h-3a2 2 0 0 0-2 2v4H2m10.5 0h1v-7a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v5.68M6 11.5h2.5"/><circle cx="4" cy="11.5" r="2"/><circle cx="10.5" cy="11.5" r="2"/></g>`),
		g.Group(children),
	)
}