package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUnlockCombinationComboKeyKeyholeLockSecureSecuritySquareUnlockUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="8" x="2" y="5.5" rx="1"/><path d="M9.47 1.53A3.44 3.44 0 0 0 7 .5A3.5 3.5 0 0 0 3.5 4v1.5"/><circle cx="7" cy="9.5" r=".5"/></g>`),
		g.Group(children),
	)
}