package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserSingleCloseGeometricHumanPersonSingleUpUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="3.75" r="3.25"/><path d="M13.18 13.5a6.49 6.49 0 0 0-12.36 0Z"/></g>`),
		g.Group(children),
	)
}