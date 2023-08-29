package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeStopWatchSixtySecondMinuteTimeStopwatchMeasureHourWholeClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5.5h3M7 .5V4m-3.89.11l1.06 1.06M1.5 8H3m.11 3.89l1.06-1.06M7 13.5V12m3.89-.11l-1.06-1.06M12.5 8H11m-.11-3.89L9.83 5.17"/>`),
		g.Group(children),
	)
}