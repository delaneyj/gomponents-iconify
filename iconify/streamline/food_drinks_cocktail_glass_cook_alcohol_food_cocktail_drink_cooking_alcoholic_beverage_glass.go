package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksCocktailGlassCookAlcoholFoodCocktailDrinkCookingAlcoholicBeverageGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.71.5a.48.48 0 0 0-.46.31a.49.49 0 0 0 .1.54L7 7l5.65-5.65a.49.49 0 0 0 .1-.54a.48.48 0 0 0-.46-.31ZM7 7v6.5m-3 0h6"/>`),
		g.Group(children),
	)
}