package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.94 13.045A9 9 0 1 0 11.987 21M3.6 9h16.8M3.6 15H13"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18a16.991 16.991 0 0 1 2.529 10.294M16 22l5-5m0 4.5V17h-4.5"/></g>`),
		g.Group(children),
	)
}