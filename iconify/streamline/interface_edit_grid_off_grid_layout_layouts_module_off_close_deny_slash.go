package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditGridOffGridLayoutLayoutsModuleOffCloseDenySlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5.5l13 13M1.79 1.79a1 1 0 0 1 .71-.29h9a1 1 0 0 1 1 1v9a1 1 0 0 1-.29.71M9 12.5H2.5a1 1 0 0 1-1-1V5m3.75-3.5v3.75m3.5-3.75v7.25m3.75-3.5H5.25m7.25 3.5H8.75m-3.5-.25v4m3.5-.5v.5M5.5 8.75h-4m.5-3.5h-.5"/>`),
		g.Group(children),
	)
}