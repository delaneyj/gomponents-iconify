package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodIceCreamOneCookFrozenBitePopsicleCreamIceCookingNutritionFreezerColdFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 10v2.5a1 1 0 0 1-1 1h0a1 1 0 0 1-1-1V10m3.75-7A1.75 1.75 0 0 1 8 1.25a1.74 1.74 0 0 1 .1-.56A3.63 3.63 0 0 0 7 .5A3.5 3.5 0 0 0 3.5 4v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V4a3.63 3.63 0 0 0-.19-1.1a1.74 1.74 0 0 1-.56.1Z"/>`),
		g.Group(children),
	)
}