package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityShieldProfileshieldSecureSecurityProfilePerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.36 13.43h0a1 1 0 0 1-.72 0h0a9.67 9.67 0 0 1-6.14-9V1.5a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2.92a9.67 9.67 0 0 1-6.14 9.01Z"/><circle cx="7" cy="5.5" r="2.5"/><path d="M3.25 11.19a5 5 0 0 1 7.5 0"/></g>`),
		g.Group(children),
	)
}