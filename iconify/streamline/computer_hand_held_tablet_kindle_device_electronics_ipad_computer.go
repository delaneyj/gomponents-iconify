package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerHandHeldTabletKindleDeviceElectronicsIpadComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="13" x="1.5" y=".5" rx="1"/><path d="M1.5 10.5h11M4.5 3h5m-5 2.5h5M4.5 8h3"/></g>`),
		g.Group(children),
	)
}