package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodIceCreamThreeCookFrozenConeCreamIceCookingNutritionFreezerColdFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.36 6.48l-2.87 6.7a.54.54 0 0 1-1 0l-2.85-6.7"/><circle cx="4.88" cy="4.75" r="2.13"/><path d="M4.88 2.63a2.12 2.12 0 1 1 4.24 0"/><circle cx="9.13" cy="4.75" r="2.13"/></g>`),
		g.Group(children),
	)
}