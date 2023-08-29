package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodIceCreamTwoCookFrozenFoodPopsicleFreezerNutritionCreamStickColdIceCooking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="9" x="3" y=".5" rx="1"/><path d="M5.5 3v3.5m3-3.5v3.5m0 3V12a1.5 1.5 0 0 1-3 0V9.5"/></g>`),
		g.Group(children),
	)
}