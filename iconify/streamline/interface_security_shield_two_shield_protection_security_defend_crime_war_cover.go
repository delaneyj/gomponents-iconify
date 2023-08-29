package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSecurityShieldTwoShieldProtectionSecurityDefendCrimeWarCover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.92 13.21l-.59.23a.94.94 0 0 1-.66 0l-.59-.23A8 8 0 0 1 1 5.78V3A6.36 6.36 0 0 0 7 .5c1.25 1.82 3.32 2.68 6 2.5v2.78a8 8 0 0 1-5.08 7.43Z"/>`),
		g.Group(children),
	)
}