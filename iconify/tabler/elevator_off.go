package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4h10a1 1 0 0 1 1 1v10m0 4a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5m7 3l2 2"/><path d="m10 14l2 2l2-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}