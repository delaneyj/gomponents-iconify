package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodBurgerDrinkBurgerFastCookCookingNutritionFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 .5h6A3.5 3.5 0 0 1 13.5 4v0a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v0A3.5 3.5 0 0 1 4 .5Zm-3.5 7h13M13 10H7l-1.5 1.5l-3-1.5H1a.5.5 0 0 0-.5.5h0a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3h0a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}