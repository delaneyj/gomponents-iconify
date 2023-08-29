package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.986 12.509a9 9 0 1 0-8.455 8.476M3.6 9h16.8M3.6 15h10.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18c2.313 3.706 3.07 7.857 2.27 12M19 16v6m3-3l-3 3l-3-3"/></g>`),
		g.Group(children),
	)
}