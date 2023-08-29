package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ufo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.95 9.01c3.02.739 5.05 2.123 5.05 3.714C22 15.091 17.52 17 12 17S2 15.091 2 12.724C2 11.134 4.04 9.739 7.07 9"/><path d="M7 9c0 1.105 2.239 2 5 2s5-.895 5-2v-.035C17 6.223 14.761 4 12 4S7 6.223 7 8.965V9m8 8l2 3m-8.5-3L7 20m5-6h.01M7 13h.01M17 13h.01"/></g>`),
		g.Group(children),
	)
}