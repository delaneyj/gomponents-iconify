package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerMonitorScreenDesktopMonitorDeviceElectronicsDisplayComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11 5.5h1.5a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H9"/><rect width="8" height="6" x=".5" y="5.5" rx="1"/><path d="M4.5 11.5v2m-2 0h4"/></g>`),
		g.Group(children),
	)
}