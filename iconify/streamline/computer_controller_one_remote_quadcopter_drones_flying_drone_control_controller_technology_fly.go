package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerControllerOneRemoteQuadcopterDronesFlyingDroneControlControllerTechnologyFly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="8" x=".5" y="5.5" rx="1"/><path d="M7 5.5v-5"/><circle cx="10" cy="9.5" r="1.5"/><circle cx="4" cy="9.5" r="1.5"/></g>`),
		g.Group(children),
	)
}