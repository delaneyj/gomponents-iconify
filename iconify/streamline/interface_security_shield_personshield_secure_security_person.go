package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityShieldPersonshieldSecureSecurityPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5" cy="2.75" r="2.25"/><path d="M6 6.61A4.49 4.49 0 0 0 .5 11v1.5H6m4.67.97h0a.5.5 0 0 1-.34 0h0A4.48 4.48 0 0 1 7.5 9.31V8a.47.47 0 0 1 .5-.5h5a.47.47 0 0 1 .5.5v1.31a4.48 4.48 0 0 1-2.83 4.16Z"/></g>`),
		g.Group(children),
	)
}