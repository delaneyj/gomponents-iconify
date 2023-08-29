package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeSleepNapSleepRestBreakClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7A6.5 6.5 0 1 1 7 .5"/><path d="M7 4.5V7L4.46 9.96M10 .5h3.5l-3.5 4h3.5"/></g>`),
		g.Group(children),
	)
}