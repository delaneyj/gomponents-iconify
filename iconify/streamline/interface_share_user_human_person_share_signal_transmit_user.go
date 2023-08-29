package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceShareUserHumanPersonShareSignalTransmitUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="6.75" r="2.25"/><path d="M11 13.5a4.5 4.5 0 0 0-8 0"/><path d="M12 10.56a6.25 6.25 0 1 0-9.92 0"/></g>`),
		g.Group(children),
	)
}