package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyTreeTwoTreePlantCircleRoundPark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 6.5L7 8v5.5M7 8l1.5-1.5"/><circle cx="7" cy="6" r="5.5"/></g>`),
		g.Group(children),
	)
}