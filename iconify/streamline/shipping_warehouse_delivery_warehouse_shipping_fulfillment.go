package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShippingWarehouseDeliveryWarehouseShippingFulfillment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 4.5h13V13a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5V4.5h0Zm0 0h0a4 4 0 0 1 4-4h5a4 4 0 0 1 4 4h0"/><path d="M11 13.5V8a.5.5 0 0 0-.5-.5h-7A.5.5 0 0 0 3 8v5.5m0-4h8m-8 2h8"/></g>`),
		g.Group(children),
	)
}