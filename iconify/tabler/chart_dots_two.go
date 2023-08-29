package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDotsTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><path d="M7 15a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4-10a2 2 0 1 0 4 0a2 2 0 1 0-4 0m5 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m5-9l-6 1.5m-.887 2.15l2.771 3.695M16 12.5l-5 2"/></g>`),
		g.Group(children),
	)
}