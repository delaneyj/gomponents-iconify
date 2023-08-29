package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodWaterMelonCookPlantPlantsDrinkCookingNutritionWatermelonFruitVegetarianFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.13 1L1 10.65a1.14 1.14 0 0 0 .48 1.55a12.22 12.22 0 0 0 11 0a1.14 1.14 0 0 0 .52-1.55L7.87 1a1 1 0 0 0-1.74 0Z"/><path d="M12.24 9.29a12 12 0 0 1-10.48 0M5.5 7.5V8m3-.5V8M7 4.5V5"/></g>`),
		g.Group(children),
	)
}