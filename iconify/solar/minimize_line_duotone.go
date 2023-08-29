package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimizeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2 22l7-7m0 0H3.143M9 15v5.857" opacity=".6"/><path d="m22 2l-7 7m0 0h5.857M15 9V3.143"/></g>`),
		g.Group(children),
	)
}