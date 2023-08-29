package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowToTopRightLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 9.5l5-5l5 5"/><path d="M12 4.5v10c0 1.667 1 5 5 5" opacity=".5"/></g>`),
		g.Group(children),
	)
}