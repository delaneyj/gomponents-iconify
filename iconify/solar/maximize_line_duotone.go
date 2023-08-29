package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 15l-7 7m0 0h5.857M2 22v-5.857" opacity=".6"/><path d="m15 9l7-7m0 0h-5.857M22 2v5.857"/></g>`),
		g.Group(children),
	)
}