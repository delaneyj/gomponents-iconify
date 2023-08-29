package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDownloadDesktopActionActionsComputerDesktopDeviceDisplayDownloadMonitorScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 3H13a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-7A.5.5 0 0 1 1 3h1.5M7 11v2.5m-2 0h4M7 .5v6"/><path d="m5 4.5l2 2l2-2"/></g>`),
		g.Group(children),
	)
}