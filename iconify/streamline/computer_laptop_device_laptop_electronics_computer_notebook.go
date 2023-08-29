package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerLaptopDeviceLaptopElectronicsComputerNotebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 8L.72 10.55a1 1 0 0 0 .05 1a1 1 0 0 0 .85.47h10.76a1 1 0 0 0 .85-.47a1 1 0 0 0 0-1L11.5 8m-8-6a1 1 0 0 0-1 1v5h9V3a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}