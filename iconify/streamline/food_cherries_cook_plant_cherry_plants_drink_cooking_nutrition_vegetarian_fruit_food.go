package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodCherriesCookPlantCherryPlantsDrinkCookingNutritionVegetarianFruitFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.36" cy="10.9" r="2.6"/><circle cx="10.64" cy="9.86" r="2.6"/><path d="M9.41 7.57A10.36 10.36 0 0 1 6 .5a13.78 13.78 0 0 0-2.6 7.8M3.36.5h5.2"/></g>`),
		g.Group(children),
	)
}