package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMedicalRibbonTwoRibbonMedicalCancerHealthBeautySymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 3.14v1.5m1.51 2.51a3.62 3.62 0 0 0 1.49-3A3.29 3.29 0 0 0 7 .64a3.28 3.28 0 0 0-3 3.5a3.62 3.62 0 0 0 1.49 3l-3.21 3.68a.51.51 0 0 0 0 .65l1.37 1.7a.47.47 0 0 0 .38.19a.52.52 0 0 0 .39-.17A23 23 0 0 0 7 9.64a23 23 0 0 0 2.6 3.55a.52.52 0 0 0 .39.17a.47.47 0 0 0 .38-.19l1.37-1.7a.5.5 0 0 0 0-.64Z"/>`),
		g.Group(children),
	)
}