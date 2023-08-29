package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMedicalSyringeInstrumentMedicalSyringeHealthBeautyNeedle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 3L2.42 7.49a2 2 0 0 0 0 2.83l1.26 1.26a2 2 0 0 0 2.83 0L11 7M10.5.5l3 3M9 5l3-3m-8.95 8.95L.5 13.5m5-12l7 7"/>`),
		g.Group(children),
	)
}