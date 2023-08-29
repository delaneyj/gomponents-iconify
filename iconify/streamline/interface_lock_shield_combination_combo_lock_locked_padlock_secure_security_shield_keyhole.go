package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLockShieldCombinationComboLockLockedPadlockSecureSecurityShieldKeyhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.36 13.4h0a1 1 0 0 1-.72 0h0A9.59 9.59 0 0 1 .5 4.46V1.54a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2.92a9.59 9.59 0 0 1-6.14 8.94Z"/><circle cx="7" cy="5.04" r="1.5"/><path d="M7 9.04v-2.5"/></g>`),
		g.Group(children),
	)
}