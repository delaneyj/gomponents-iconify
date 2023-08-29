package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceComputerCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 10a7 7 0 1 0 14 0a7 7 0 1 0-14 0"/><path d="M9 10a3 3 0 1 0 6 0a3 3 0 1 0-6 0m-1 6l-2.091 3.486A1 1 0 0 0 6.766 21h10.468a1 1 0 0 0 .857-1.514L16 16"/></g>`),
		g.Group(children),
	)
}