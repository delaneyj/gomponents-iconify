package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelLaundryBasketBathroomLaundryBinBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 1v9a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3V1M.5 1h13m-12 3h3m-3 3h3m-3 3h3m5-6h3m-3 3h3m-3 3h3"/>`),
		g.Group(children),
	)
}