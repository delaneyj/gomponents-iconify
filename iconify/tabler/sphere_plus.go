package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpherePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12c0 1.657 4.03 3 9 3c1.116 0 2.185-.068 3.172-.192m5.724-2.35A1.1 1.1 0 0 0 21 12"/><path d="M20.984 12.546a9 9 0 1 0-8.442 8.438M16 19h6m-3-3v6"/></g>`),
		g.Group(children),
	)
}