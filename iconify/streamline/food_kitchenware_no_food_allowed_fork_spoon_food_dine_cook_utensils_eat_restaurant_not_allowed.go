package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodKitchenwareNoFoodAllowedForkSpoonFoodDineCookUtensilsEatRestaurantNotAllowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5.5v10M6 .5V3a2.5 2.5 0 0 1-2.5 2.5h0A2.5 2.5 0 0 1 1 3V.5m-.5 13l13-13m-1 5A14.61 14.61 0 0 1 13 10H8.5m0-4.5v-5a7.41 7.41 0 0 1 2.66 2.34M8.5 13.5V9"/>`),
		g.Group(children),
	)
}