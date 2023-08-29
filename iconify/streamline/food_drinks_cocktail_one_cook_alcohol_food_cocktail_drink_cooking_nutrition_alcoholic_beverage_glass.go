package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksCocktailOneCookAlcoholFoodCocktailDrinkCookingNutritionAlcoholicBeverageGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 8.5L2.48 1.26A.5.5 0 0 1 2.9.5h8.2a.5.5 0 0 1 .42.76Zm0 0v5m-2.5 0h5M3.56 3h6.88"/>`),
		g.Group(children),
	)
}