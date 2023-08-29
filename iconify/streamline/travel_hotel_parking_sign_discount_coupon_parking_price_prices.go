package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelParkingSignDiscountCouponParkingPricePrices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="3"/><path d="M5.5 3.5v7m0-3h2a2 2 0 0 0 0-4h-2"/></g>`),
		g.Group(children),
	)
}