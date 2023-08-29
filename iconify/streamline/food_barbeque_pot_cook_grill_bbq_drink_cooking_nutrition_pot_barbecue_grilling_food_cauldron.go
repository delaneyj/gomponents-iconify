package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodBarbequePotCookGrillBbqDrinkCookingNutritionPotBarbecueGrillingFoodCauldron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 12a6.51 6.51 0 0 0 6.48-6a.5.5 0 0 0-.48-.5H1a.5.5 0 0 0-.48.5A6.51 6.51 0 0 0 7 12Zm-2.95-.71L3.5 13.5m6.45-2.21l.55 2.21M3.5 3V2m7 1V2M7 3V.5"/>`),
		g.Group(children),
	)
}