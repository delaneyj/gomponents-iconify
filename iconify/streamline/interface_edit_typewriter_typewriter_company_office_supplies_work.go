package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditTypewriterTypewriterCompanyOfficeSuppliesWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 3.5h13v3H.5zM3 .5h8a.5.5 0 0 1 .5.5v2.5h0h-9h0V1A.5.5 0 0 1 3 .5Z"/><circle cx="3" cy="8" r="1.5"/><circle cx="11" cy="8" r="1.5"/><path d="M2.5 9.41v3.09a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V9.41m-9 .09h9"/></g>`),
		g.Group(children),
	)
}