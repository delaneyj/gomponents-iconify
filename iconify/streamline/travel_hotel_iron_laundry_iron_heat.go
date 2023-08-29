package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelIronLaundryIronHeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.59 5.5H8.91a8.78 8.78 0 0 0-8.32 6h11.82a1 1 0 0 0 .76-.35a1 1 0 0 0 .23-.81Zm0 0l-.22-1.33a2 2 0 0 0-2-1.67H4.59"/>`),
		g.Group(children),
	)
}