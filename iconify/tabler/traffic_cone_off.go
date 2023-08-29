package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficConeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16M9.4 10h.6m4 0h.6m-6.8 5H15m-9 5L9.5 9.5m1-3L11 5h2l2 6m2 6l1 3M3 3l18 18"/>`),
		g.Group(children),
	)
}