package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleeperChair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 18v3m1-11V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5"/><path d="M19.5 10a2.5 2.5 0 0 0-2.5 2.5V14H7v-1.5a2.5 2.5 0 1 0-3 2.45V18h16v-3.05a2.5 2.5 0 0 0-.5-4.95Zm.5 8v3"/></g>`),
		g.Group(children),
	)
}