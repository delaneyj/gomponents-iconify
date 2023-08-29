package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeStopWatchHalfQuarterSecondsThirtyHourHalfMeasureTimeStopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.75 2.5a5.5 5.5 0 0 1 0 11Zm-1.5-2h3m-1.5 0v2m5-.5l1 1M2.86 4.11l1.06 1.06M1.25 8h1.5m.11 3.89l1.06-1.06"/>`),
		g.Group(children),
	)
}