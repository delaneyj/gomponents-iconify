package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerChipTwoCoreMicroprocessorDeviceElectronicsChipComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="13" x="3" y=".5" rx="1"/><path d="M3 3.5H.5M3 7H.5M3 10.5H.5m13-7H11M13.5 7H11m2.5 3.5H11m-4.5 0h2"/></g>`),
		g.Group(children),
	)
}