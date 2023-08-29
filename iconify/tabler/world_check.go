package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.946 12.99a9 9 0 1 0-9.46 7.995M3.6 9h16.8M3.6 15h13.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18a16.997 16.997 0 0 1 2.311 12.001M15 19l2 2l4-4"/></g>`),
		g.Group(children),
	)
}