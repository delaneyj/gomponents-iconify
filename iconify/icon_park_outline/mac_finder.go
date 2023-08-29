package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacFinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 38V10a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2Z"/><path d="M25 8s-5 10-4 17h6l1 15m6 0H22m8-32H18m16 8v2m-20-2v2"/><path d="M13 29s4.19 3 11 3s11-3 11-3"/></g>`),
		g.Group(children),
	)
}