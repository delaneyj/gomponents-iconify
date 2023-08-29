package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPrinterPrinterCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 8.5h9a1 1 0 0 1 1 1v4h0h-11h0v-4a1 1 0 0 1 1-1Z"/><path d="M1.5 9.5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1M3 .5h8v5H3zM4 11h6"/></g>`),
		g.Group(children),
	)
}