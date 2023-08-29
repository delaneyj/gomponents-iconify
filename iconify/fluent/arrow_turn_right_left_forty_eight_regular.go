package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRightLeftFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.526 9.407a1.25 1.25 0 0 1 .948-2.313l31.31 12.824c1.864.763 1.864 3.401 0 4.164L11.578 36.455l8.713 4.168a1.25 1.25 0 1 1-1.08 2.255l-11.5-5.5a1.25 1.25 0 0 1-.614-1.608l5-12a1.25 1.25 0 1 1 2.308.961l-3.951 9.483L40.273 22L9.525 9.407Z"/>`),
		g.Group(children),
	)
}