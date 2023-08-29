package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 5v4"/><path d="M5 19V9a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2Z"/><path d="M5 10V5a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v5"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.167 11L9.5 14h4l-1.667 3"/></g>`),
		g.Group(children),
	)
}