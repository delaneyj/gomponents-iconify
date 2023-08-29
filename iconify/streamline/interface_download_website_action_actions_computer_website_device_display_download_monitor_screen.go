package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDownloadWebsiteActionActionsComputerWebsiteDeviceDisplayDownloadMonitorScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 7v6.5M4.5 11L7 13.5L9.5 11"/><path d="M2.5 13.5h-1a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-1m-11-10h13"/></g>`),
		g.Group(children),
	)
}