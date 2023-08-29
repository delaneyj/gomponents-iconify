package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMedicalHeartRateHealthBeautyInformationDataBeatPulseMonitorHeartRateInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.58 4.31C1.09 1.85 4.12 0 7 3.27c4.11-4.71 8.5 1.13 5.52 4.14L7 12.5l-3.23-3"/><path d="M.5 7H3l1.5-2l2 3.5l1.5-2h1.5"/></g>`),
		g.Group(children),
	)
}