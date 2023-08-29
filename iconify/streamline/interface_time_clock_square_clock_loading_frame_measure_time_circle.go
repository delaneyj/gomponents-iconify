package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeClockSquareClockLoadingFrameMeasureTimeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="4"/><path d="M7 5.5V7h1.5"/><rect width="13" height="13" x=".5" y=".5" rx="1"/></g>`),
		g.Group(children),
	)
}