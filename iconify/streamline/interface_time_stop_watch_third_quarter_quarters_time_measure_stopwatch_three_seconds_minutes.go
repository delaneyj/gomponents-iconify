package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeStopWatchThirdQuarterQuartersTimeMeasureStopwatchThreeSecondsMinutes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.75 2.5A5.5 5.5 0 1 1 1.25 8h5.5Zm-1.5-2h3m-1.5 2v-2m5 1.5l1 1M2.86 4.11l1.06 1.06"/>`),
		g.Group(children),
	)
}