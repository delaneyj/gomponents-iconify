package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeResetTimeClockResetStopwatchCircleMeasureLoading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 .5A6.5 6.5 0 1 1 .5 7a7.23 7.23 0 0 1 2-5"/><path d="m.5 2.5l2-.5l.5 2m4-.5v4l2.6 1.3"/></g>`),
		g.Group(children),
	)
}