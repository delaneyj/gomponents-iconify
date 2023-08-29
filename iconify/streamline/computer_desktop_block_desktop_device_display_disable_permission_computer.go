package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopBlockDesktopDeviceDisplayDisablePermissionComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 2H13a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6"/><circle cx="7" cy="5" r="3"/><path d="m4.88 7.12l4.24-4.24"/></g>`),
		g.Group(children),
	)
}