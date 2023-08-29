package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHundredSixty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 15.328c2.414-.718 4-1.94 4-3.328c0-2.21-4.03-4-9-4s-9 1.79-9 4s4.03 4 9 4"/><path d="m9 13l3 3l-3 3"/></g>`),
		g.Group(children),
	)
}