package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodKitchenwareRefrigeratorFridgeCookAppliancesCookingNutritionFreezerApplianceFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="12" x="2.5" y=".5" rx="1"/><path d="M2.5 6h9m-7-3v.5m-.5 9v1m6-1v1"/></g>`),
		g.Group(children),
	)
}