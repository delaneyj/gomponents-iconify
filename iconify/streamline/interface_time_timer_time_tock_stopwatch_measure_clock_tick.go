package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeTimerTimeTockStopwatchMeasureClockTick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="8" r="5.5"/><path d="M.5 2.5A8.69 8.69 0 0 1 3 .5m10.5 2a8.69 8.69 0 0 0-2.5-2M7 5v3h2.5"/></g>`),
		g.Group(children),
	)
}