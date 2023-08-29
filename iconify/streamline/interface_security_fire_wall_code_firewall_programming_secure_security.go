package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityFireWallCodeFirewallProgrammingSecureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 7.5h13v6H.5z"/><path d="M7 .5S8.5 3 6.5 4C5.5 4.5 4 3 4 3a3.18 3.18 0 0 0 3 4.5A3.34 3.34 0 0 0 10.5 4A3.34 3.34 0 0 0 7 .5Zm-6.5 10h13m-8 0v3m3-6v3"/></g>`),
		g.Group(children),
	)
}