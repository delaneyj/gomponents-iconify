package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyBed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 18a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-8 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0M4 8l2.1 2.8A3 3 0 0 0 8.5 12H20M10 6h4m-2-2v4"/><path d="M12 12v2l-2.5 2.5m5 0L12 14"/></g>`),
		g.Group(children),
	)
}