package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficLightsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4c.912-1.219 2.36-2 4-2a5 5 0 0 1 5 5v6m0 4a5 5 0 0 1-10 0V7"/><path d="M12 8a1 1 0 1 0-1-1m.291 4.295a1 1 0 0 0 1.418 1.41M11 17a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}