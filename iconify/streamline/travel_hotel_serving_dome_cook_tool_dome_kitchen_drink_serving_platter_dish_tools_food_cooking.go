package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelServingDomeCookToolDomeKitchenDrinkServingPlatterDishToolsFoodCooking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 9.75v-.5a5.5 5.5 0 0 0-11 0v.5m12 0H.5l.32 1.07a2 2 0 0 0 1.92 1.43h8.52a2 2 0 0 0 1.92-1.43Zm-6.5-6v-2"/>`),
		g.Group(children),
	)
}