package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandSurveying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 26v18m0-18l12 18M24 26L12 44m12-30v6m-5 0h10"/><path d="M10 6h26v8H10z"/><path stroke-linecap="round" d="M40 8v4"/></g>`),
		g.Group(children),
	)
}