package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksCocktailShakerAlcoholDrinkMixShakeCocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 6l-.89 6.62a1 1 0 0 1-1 .88H4.38a1 1 0 0 1-1-.88L2.5 6m8.23-2.32a1 1 0 0 0-1-.68H4.22a1 1 0 0 0-.95.68L2.5 6h9Z"/><path d="M8.5 3V1.5a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1V3"/></g>`),
		g.Group(children),
	)
}