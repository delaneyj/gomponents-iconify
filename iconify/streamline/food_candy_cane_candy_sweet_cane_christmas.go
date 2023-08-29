package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodCandyCaneCandySweetCaneChristmas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5a1.5 1.5 0 0 1 3 0v7a1.5 1.5 0 0 0 3 0V5a4.5 4.5 0 0 0-9 0v1a1.5 1.5 0 0 0 3 0Zm3 5l3-3M8.42 4.51L9.95 1.6M5.86 4.02L3.03 2.88"/>`),
		g.Group(children),
	)
}