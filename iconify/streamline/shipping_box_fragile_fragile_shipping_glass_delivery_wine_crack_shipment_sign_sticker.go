package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShippingBoxFragileFragileShippingGlassDeliveryWineCrackShipmentSignSticker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.67.5a1 1 0 0 0-1 .89C2.56 2.45 2.5 4.05 2.5 4.5c0 2.62 2 4 4.5 4s4.5-1.38 4.5-4c0-.45-.06-2-.17-3.11a1 1 0 0 0-1-.89ZM7 8.5v5m-2 0h4"/><path d="m7 .5l-1.5 2l2 1.5l-1 1.5"/></g>`),
		g.Group(children),
	)
}