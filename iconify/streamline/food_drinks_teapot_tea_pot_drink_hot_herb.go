package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksTeapotTeaPotDrinkHotHerb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8 6a5.51 5.51 0 0 0-4.46 2.28L1.5 7l-1 1l2 3.5v1a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-1A5.5 5.5 0 0 0 8 6Zm0 0V4.5"/><path d="M4 7.73V4.5a4 4 0 0 1 8 0v3.23"/></g>`),
		g.Group(children),
	)
}