package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrumStickOneCookAnimalDrumsticksProductsChickenCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.49 9.6a1.66 1.66 0 0 0-3-.79l-1-1s-.08-.05-.11-.08A5.6 5.6 0 0 0 8 1.87a4.15 4.15 0 0 0-6.13 0A4.16 4.16 0 0 0 1.83 8a5.58 5.58 0 0 0 6 1.34a.56.56 0 0 0 0 .08L9 10.58A1.63 1.63 0 0 0 8.43 12a1.66 1.66 0 0 0 3.31-.34a1.6 1.6 0 0 0-.05-.21H12a1.65 1.65 0 0 0 1.49-1.85Z"/>`),
		g.Group(children),
	)
}