package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksWineBottleCookBottleWineDrinkCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.25 4V1a.5.5 0 0 0-.5-.5h-1.5a.5.5 0 0 0-.5.5v3h0a3.36 3.36 0 0 0-1.5 2.8v5.7a1 1 0 0 0 1 1h3.5a1 1 0 0 0 1-1V6.8A3.36 3.36 0 0 0 8.25 4Zm-4 3.5h5.5m0 3.5h-5.5"/>`),
		g.Group(children),
	)
}