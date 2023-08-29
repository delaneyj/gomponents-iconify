package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopAddDesktopDeviceDisplayAddPlusComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="2" rx=".5"/><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6m-.5-7h-5M7 4v5"/></g>`),
		g.Group(children),
	)
}