package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityShieldFourShieldProtectionSecurityDefendCrimeWarCover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.36 13.43h0a1 1 0 0 1-.72 0h0A9.57 9.57 0 0 1 .5 4.49V1.88a.51.51 0 0 1 .36-.49a21.57 21.57 0 0 1 12.28 0a.51.51 0 0 1 .36.49v2.61a9.57 9.57 0 0 1-6.14 8.94Z"/>`),
		g.Group(children),
	)
}