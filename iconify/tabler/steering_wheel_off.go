package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteeringWheelOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.04 16.048A9 9 0 0 0 7.957 3.958m-2.32 1.678a9 9 0 1 0 12.737 12.719"/><path d="M10.595 10.576a2 2 0 1 0 2.827 2.83M12 14v7m-2-9l-6.75-2m12.292 1.543L20.75 10M3 3l18 18"/></g>`),
		g.Group(children),
	)
}