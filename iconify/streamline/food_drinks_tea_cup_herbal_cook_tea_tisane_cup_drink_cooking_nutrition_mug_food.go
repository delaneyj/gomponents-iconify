package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksTeaCupHerbalCookTeaTisaneCupDrinkCookingNutritionMugFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 4h8a.5.5 0 0 1 .5.5V9a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2V4.5A.5.5 0 0 1 2 4ZM.5 13.5h13m-3-8.5h1a2 2 0 0 1 2 2h0a2 2 0 0 1-2 2h-1M3 .5v1m5-1v1M5.5.5v1"/>`),
		g.Group(children),
	)
}