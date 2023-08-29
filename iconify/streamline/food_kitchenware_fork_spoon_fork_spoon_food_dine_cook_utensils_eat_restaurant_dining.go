package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodKitchenwareForkSpoonForkSpoonFoodDineCookUtensilsEatRestaurantDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="10.6" cy="3.5" rx="2.4" ry="3"/><path d="M10.6 6.5v7M3.5.5v13M6 .5V3a2.5 2.5 0 0 1-2.5 2.5h0A2.5 2.5 0 0 1 1 3V.5"/></g>`),
		g.Group(children),
	)
}