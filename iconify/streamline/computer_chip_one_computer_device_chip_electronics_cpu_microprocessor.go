package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerChipOneComputerDeviceChipElectronicsCpuMicroprocessor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="8" x="3" y="3" rx="1"/><path d="M5 3V.5M9 3V.5M3 9H.5M3 5H.5M9 11v2.5M5 11v2.5M11 5h2.5M11 9h2.5m-5-1.5h-2"/></g>`),
		g.Group(children),
	)
}