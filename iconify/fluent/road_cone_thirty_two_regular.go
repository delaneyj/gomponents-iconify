package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadConeThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.77 2a1.5 1.5 0 0 0-1.434 1.064L5.758 28H3a1 1 0 1 0 0 2h26a1 1 0 1 0 0-2h-2.759L18.665 3.064A1.5 1.5 0 0 0 17.229 2h-2.458ZM7.85 28l1.66-5.462a.996.996 0 0 0 .491.129h8.5a1 1 0 1 0 0-2h-8.423l.81-2.667H17a1 1 0 1 0 0-2h-5.505l3.646-12h1.718l7.292 24H7.849Z"/>`),
		g.Group(children),
	)
}