package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodBaguetteCookBreadGlutenDrinkCookingNutritionBaguetteFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.636 6.93l6.468-5.074a3 3 0 0 1 4.212.509l.975 1.243a1 1 0 0 1-.17 1.404l-9.615 7.542a1 1 0 0 1-1.404-.17l-.975-1.243a3 3 0 0 1 .51-4.212Zm1.234-.97l1.56 2m2.03-4.82l1.57 2"/>`),
		g.Group(children),
	)
}