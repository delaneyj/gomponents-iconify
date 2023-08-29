package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GgCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m717 1354l271-271l-279-279l-88 88l192 191l-96 96l-279-279l279-279l40 40l87-87l-127-128l-454 454zm358-8l454-454l-454-454l-271 271l279 279l88-88l-192-191l96-96l279 279l-279 279l-40-40l-87 88zm717-450q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348z"/>`),
		g.Group(children),
	)
}