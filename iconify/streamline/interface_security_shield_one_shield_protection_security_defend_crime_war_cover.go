package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityShieldOneShieldProtectionSecurityDefendCrimeWarCover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.36 13.43h0a1 1 0 0 1-.72 0h0A9.57 9.57 0 0 1 .5 4.49v-3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v3a9.57 9.57 0 0 1-6.14 8.94Z"/>`),
		g.Group(children),
	)
}