package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMedicalRibbonOneRibbonMedicalCancerHealthBeautySymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 13.5S10 6.39 10 3.5a2.84 2.84 0 0 0-3-3a2.84 2.84 0 0 0-3 3c0 2.89 6.5 10 6.5 10"/>`),
		g.Group(children),
	)
}