package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditGlueGlueCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 8v4.5a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1V8A2.5 2.5 0 0 1 5 5.5h4A2.5 2.5 0 0 1 11.5 8ZM7.5.5h-1L5 5.5h4L7.5.5zM5 11h4M2.5 8.5h9"/>`),
		g.Group(children),
	)
}