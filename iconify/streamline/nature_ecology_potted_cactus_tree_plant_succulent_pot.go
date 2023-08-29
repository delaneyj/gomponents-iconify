package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyPottedCactusTreePlantSucculentPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.1 12.61a1 1 0 0 1-1 .89H4.9a1 1 0 0 1-1-.89L3.5 9h7ZM4.5 9V5a2.5 2.5 0 0 1 5 0v4M7 2.5v-2M4.5 6h-2m1-3.5l1.31 1.31M9.5 6h2m-1-3.5L9.19 3.81"/>`),
		g.Group(children),
	)
}