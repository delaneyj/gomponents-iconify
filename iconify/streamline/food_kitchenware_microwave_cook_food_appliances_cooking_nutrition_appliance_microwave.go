package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodKitchenwareMicrowaveCookFoodAppliancesCookingNutritionApplianceMicrowave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2.5" rx=".5"/><rect width="6.5" height="5" x="2.5" y="4.5" rx=".5"/><path d="M11 5h.5M11 6.5h.5M11.25 9h0"/></g>`),
		g.Group(children),
	)
}