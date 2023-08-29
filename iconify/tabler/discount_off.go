package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscountOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 15l3-3m2-2l1-1m-5.852.145A.498.498 0 0 0 9.5 10a.5.5 0 0 0 .35-.142m4.298 4.287A.498.498 0 0 0 14.5 15a.5.5 0 0 0 .35-.142"/><path d="M5.641 5.631A9 9 0 1 0 18.36 18.369m1.68-2.318A9 9 0 0 0 7.966 3.953M3 3l18 18"/></g>`),
		g.Group(children),
	)
}