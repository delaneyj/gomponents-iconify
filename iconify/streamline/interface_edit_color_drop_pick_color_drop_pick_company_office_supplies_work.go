package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditColorDropPickColorDropPickCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 9C12 5.5 7 .5 7 .5S2 5.5 2 9a4.77 4.77 0 0 0 5 4.5A4.77 4.77 0 0 0 12 9ZM7 .5v13m2.5-.6V3.44"/>`),
		g.Group(children),
	)
}