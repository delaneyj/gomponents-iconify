package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShippingTransferTruckTimeTruckShippingDeliveryTimeWaitingDelay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6" cy="12.49" r="1"/><circle cx="10.5" cy="12.49" r="1"/><path d="M8.38 3.53A4 4 0 1 0 2 7.62m2.5-3.11L6 3.01m.5 6.49v-1H5a2 2 0 0 0-2 2v2"/><path d="M13.5 12.49v-5a1 1 0 0 0-1-1h-5a1 1 0 0 0-1 1v2"/></g>`),
		g.Group(children),
	)
}