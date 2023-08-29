package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 5l7-3l6 3l6.303-2.701a.5.5 0 0 1 .697.46V19l-7 3l-6-3l-6.303 2.701a.5.5 0 0 1-.697-.46V5Zm14 14.395l4-1.714V5.033l-4 1.714v12.648Zm-2-.131V6.736l-4-2v12.528l4 2Zm-6-2.011V4.605L4 6.319v12.648l4-1.714Z"/>`),
		g.Group(children),
	)
}