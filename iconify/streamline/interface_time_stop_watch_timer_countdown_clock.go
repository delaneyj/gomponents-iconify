package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeStopWatchTimerCountdownClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 6.92a7.62 7.62 0 0 1-.74 3.21M7 13.42a6.5 6.5 0 0 1-6.5-6.5m8.5 2l-2-2.5H4"/><circle cx="12" cy="11.42" r="1.5"/><path d="M.92 4.52A6.31 6.31 0 0 1 1.79 3m1.93-1.72a6.3 6.3 0 0 1 1.81-.7m7.55 3.94A6.31 6.31 0 0 0 12.21 3m-1.93-1.72a6.3 6.3 0 0 0-1.81-.7"/></g>`),
		g.Group(children),
	)
}