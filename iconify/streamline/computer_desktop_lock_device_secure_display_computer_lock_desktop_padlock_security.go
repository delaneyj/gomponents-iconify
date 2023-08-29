package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopLockDeviceSecureDisplayComputerLockDesktopPadlockSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 2H13a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h2.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6"/><path d="M5 5.5h4v3H5zm.5 0v-1a1.5 1.5 0 0 1 3 0v1"/></g>`),
		g.Group(children),
	)
}