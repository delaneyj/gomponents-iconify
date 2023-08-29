package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VegetableBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 14H7.378a3 3 0 0 0-2.98 3.354L7.12 40.236A2 2 0 0 0 9.105 42h30.368a2 2 0 0 0 1.991-1.807l2.218-22.904A3 3 0 0 0 40.696 14H38M5 22h38m-28 7h18m-16 7h14"/><path d="M24 6c-4.418 0-8 6.925-8 15.467c0 .178.002.356.005.533h15.99c.003-.177.005-.355.005-.533C32 12.925 28.418 6 24 6Z"/></g>`),
		g.Group(children),
	)
}