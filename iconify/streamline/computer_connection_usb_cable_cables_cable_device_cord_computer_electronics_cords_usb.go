package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerConnectionUsbCableCablesCableDeviceCordComputerElectronicsCordsUsb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V11M4.5 4V1.5a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1V4m1 0h-7a.5.5 0 0 0-.5.5v2.59a1 1 0 0 0 .29.7L4.5 9v1a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V9l1.21-1.21a1 1 0 0 0 .29-.7V4.5a.5.5 0 0 0-.5-.5Zm-4 2.5h1"/>`),
		g.Group(children),
	)
}