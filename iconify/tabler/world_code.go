package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.942 13.02a9 9 0 1 0-9.47 7.964M3.6 9h16.8M3.6 15h9.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18c2 3.206 2.837 6.913 2.508 10.537M20 21l2-2l-2-2m-3 0l-2 2l2 2"/></g>`),
		g.Group(children),
	)
}