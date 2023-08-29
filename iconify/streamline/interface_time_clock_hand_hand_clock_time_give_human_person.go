package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeClockHandHandClockTimeGiveHumanPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.28 5.25a4.5 4.5 0 1 1 7.93 3.38"/><path d="M8.75 3.75v2h1.5m-9.5 2h4a2.5 2.5 0 0 1 2.5 2.5h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5h-10"/></g>`),
		g.Group(children),
	)
}