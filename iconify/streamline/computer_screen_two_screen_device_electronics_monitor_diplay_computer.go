package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerScreenTwoScreenDeviceElectronicsMonitorDiplayComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="1.5" rx=".5"/><path d="M3.53 13.5a3.51 3.51 0 0 1 6.94 0"/></g>`),
		g.Group(children),
	)
}