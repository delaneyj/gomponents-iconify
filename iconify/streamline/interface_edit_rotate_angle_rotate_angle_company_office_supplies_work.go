package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditRotateAngleRotateAngleCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v13h13"/><path d="M7.5 13.5a7 7 0 0 0-7-7m0 7L3 11m2.5-2.5L7 7m2-2l1.5-1.5m2-2l1-1"/></g>`),
		g.Group(children),
	)
}