package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsPlaneTakeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M3 18h18v2H3zm18.509-9.473a1.61 1.61 0 0 0-2.036-1.019L15 9L7 6L5 7l6 4l-4 2l-4-2l-1 1l4 4l14.547-5.455a1.611 1.611 0 0 0 .962-2.018z" fill="currentColor"/>`),
		g.Group(children),
	)
}