package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceShareHandLockGiveHandLockPadlockSecureSecurityTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6" height="4.5" x="7.5" y="3.5" rx=".5"/><path d="M8.5 3.5v-1a2 2 0 0 1 4 0v1m-12 5h3A2.5 2.5 0 0 1 6 11h0m-3 0h5a2 2 0 0 1 2 2h0a.5.5 0 0 1-.5.5h-9"/></g>`),
		g.Group(children),
	)
}