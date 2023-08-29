package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilitaryAward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 13a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/><path d="M8.5 10.5L7.5 8H2l2.48 5.788A2 2 0 0 0 6.32 15H8.5m7-4.5l1-2.5H22l-2.48 5.788A2 2 0 0 1 17.68 15H15.5"/></g>`),
		g.Group(children),
	)
}