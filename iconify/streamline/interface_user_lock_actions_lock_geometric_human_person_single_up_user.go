package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserLockActionsLockGeometricHumanPersonSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M6 6.61A4.49 4.49 0 0 0 .5 11v1.5h4"/><rect width="6" height="4.5" x="7.5" y="9" rx=".5"/><path d="M8.5 9V8a2 2 0 0 1 4 0v1"/></g>`),
		g.Group(children),
	)
}