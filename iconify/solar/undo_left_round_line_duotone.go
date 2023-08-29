package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoLeftRoundLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7 4L4 7l3 3"/><path d="M4 7h11a5 5 0 0 1 0 10H8" opacity=".5"/></g>`),
		g.Group(children),
	)
}