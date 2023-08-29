package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodDrinksWaterGlassGlassWaterJuiceDrinkLiquid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.2 11.73a2 2 0 0 1-2 1.77H4.78a2 2 0 0 1-2-1.77L1.5.5h11ZM1.96 4.5h10.08"/>`),
		g.Group(children),
	)
}