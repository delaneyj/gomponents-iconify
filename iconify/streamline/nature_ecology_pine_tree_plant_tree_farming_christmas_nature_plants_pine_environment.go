package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyPineTreePlantTreeFarmingChristmasNaturePlantsPineEnvironment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 10.5H2l3.5-5H3l4-5l4 5H8.5l3.5 5zm-5 0v3"/>`),
		g.Group(children),
	)
}