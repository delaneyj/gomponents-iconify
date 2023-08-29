package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerPrinterScanDeviceElectronicsPrinterPrintComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 11h2a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1V10a1 1 0 0 0 1 1h2"/><path d="M3.5 9.5V13a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5V9.5Zm7-5V1a.5.5 0 0 0-.5-.5H4a.5.5 0 0 0-.5.5v3.5M11 7h-1"/></g>`),
		g.Group(children),
	)
}