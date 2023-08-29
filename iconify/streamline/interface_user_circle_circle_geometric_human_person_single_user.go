package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserCircleCircleGeometricHumanPersonSingleUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2.5"/><path d="M2.73 11.9a5 5 0 0 1 8.54 0"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}