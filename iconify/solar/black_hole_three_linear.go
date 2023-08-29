package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackHoleThreeLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="2"/><path stroke-linecap="round" d="M12 10c5 0 4.6 12-3 12"/><path stroke-linecap="round" d="M12.312 14c-5 0-4.6-12 3-12"/><path stroke-linecap="round" d="M10 12.312c0-5 12-4.6 12 3"/><path stroke-linecap="round" d="M14 12c0 5-12 4.6-12-3"/></g>`),
		g.Group(children),
	)
}