package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsJoinTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7h1.948c1.913 0 3.705.933 4.802 2.5a5.861 5.861 0 0 0 4.802 2.5H21"/><path d="M3 17h1.95a5.854 5.854 0 0 0 4.798-2.5a5.854 5.854 0 0 1 4.798-2.5H20"/><path d="m18 15l3-3l-3-3"/></g>`),
		g.Group(children),
	)
}