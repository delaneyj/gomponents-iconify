package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksMilkCanisterMilkDairyCanisterDrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11 7.5l-1.5-2h-5L3 7.5v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1Zm0 0l2.5-2M3 7.5l-2.5-2m4-5h5a1 1 0 0 1 1 1V3h0h-7h0V1.5a1 1 0 0 1 1-1Zm0 2.5v2.5m5-2.5v2.5"/>`),
		g.Group(children),
	)
}