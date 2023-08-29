package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodCheeseOneCookCheeseAnimalProductsCookingNutritionDairyFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.59 5.75H13M9.09.77L9 .75L.75 5.6A.5.5 0 0 0 .5 6v1.75c1 0 2.5 0 2.5 1.75S1.47 11.25.5 11.25v1.5a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-7A5.88 5.88 0 0 0 9.09.77Z"/><circle cx="10" cy="8.25" r=".5"/><circle cx="6.5" cy="10.25" r=".5"/></g>`),
		g.Group(children),
	)
}