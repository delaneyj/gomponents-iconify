package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserBlockActionsBlockCloseDeniedDenyGeometricHumanPersonSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><circle cx="10.25" cy="10.25" r="3.25"/><path d="m7.95 12.55l4.6-4.6M6 6.61A4.49 4.49 0 0 0 .5 11v1.5h4"/></g>`),
		g.Group(children),
	)
}