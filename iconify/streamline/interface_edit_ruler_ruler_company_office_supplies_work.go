package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditRulerRulerCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.502 9.496L9.503.495l4.002 4.002l-9.001 9.001zM7.5 2.5L9 4M5 5l1.5 1.5m-4 1L4 9"/>`),
		g.Group(children),
	)
}