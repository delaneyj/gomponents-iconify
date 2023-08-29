package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodPizzaDrinkCookFastCookingNutritionPizzaFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.75.5L8 6l5.5-1.75A3.75 3.75 0 0 0 9.75.5Z"/><path d="M7.5.52L7 .5A6.5 6.5 0 1 0 13.5 7v-.5"/><circle cx="5.5" cy="9.5" r=".5"/><circle cx="4" cy="5" r=".5"/><circle cx="10.5" cy="8" r=".5"/></g>`),
		g.Group(children),
	)
}