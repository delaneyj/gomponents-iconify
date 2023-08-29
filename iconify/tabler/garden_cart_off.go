package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GardenCartOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.733 15.732a2.5 2.5 0 1 0 3.544 3.527M6 8v11a1 1 0 0 0 1.806.591L11.5 14.5v.055"/><path d="M6 8h2m4 0h9l-3 6.01m-3.319.693l-4.276-.45a4 4 0 0 1-3.296-2.493L4.256 4.63A1 1 0 0 0 3.328 4H2.005M3 3l18 18"/></g>`),
		g.Group(children),
	)
}