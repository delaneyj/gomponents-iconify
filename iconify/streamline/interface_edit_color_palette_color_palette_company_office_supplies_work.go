package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditColorPaletteColorPaletteCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.5" cy="9" r="4"/><circle cx="9.5" cy="9" r="4"/><circle cx="7" cy="5" r="4"/></g>`),
		g.Group(children),
	)
}