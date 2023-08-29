package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditCutterCutterCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 7.5l-2 2v4l4-4"/><path d="M11.67 1.33a2.82 2.82 0 0 0-4 0L2.5 6.5l4 4l5.17-5.17a2.82 2.82 0 0 0 0-4ZM6.5 6.5l3-3"/></g>`),
		g.Group(children),
	)
}