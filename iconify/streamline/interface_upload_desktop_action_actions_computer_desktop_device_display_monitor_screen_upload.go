package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUploadDesktopActionActionsComputerDesktopDeviceDisplayMonitorScreenUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 3h1a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-7A.5.5 0 0 1 1 3h1m5 8v2.5m-2 0h4m-2-7v-6m-2 2l2-2l2 2"/>`),
		g.Group(children),
	)
}