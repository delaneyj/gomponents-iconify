package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksBeerMugBeerCookBreweryDrinkMugCookingNutritionBrewBrewingFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5.5h8a1 1 0 0 1 1 1v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a1 1 0 0 1 1-1Zm9 2.5h2a1 1 0 0 1 1 1v4.5a1 1 0 0 1-1 1h-2M3.75 4v5.5M7.25 4v5.5"/>`),
		g.Group(children),
	)
}