package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodKitchenwareSpoonPlateForkPlateFoodDineCookUtensilsEatRestaurantDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v2.25a2.5 2.5 0 0 0 5 0V.5M3 .5v13m10.5 0a6.5 6.5 0 0 1 0-13"/><path d="M13.5 10.5a3.5 3.5 0 0 1 0-7"/></g>`),
		g.Group(children),
	)
}