package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeStopWatchAlternateTimerCountdownClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5A6.5 6.5 0 1 1 13.5 7a7.23 7.23 0 0 1-2 5"/><path d="m13.5 11.5l-2 .5l-.5-2M9 9L7 6.5H4"/></g>`),
		g.Group(children),
	)
}