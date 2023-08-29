package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisXArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 20.5L3 15.03l1.46 2.57L11 13.82V3h2v10.82l9.39 5.43l-1 1.75L12 15.56l-6.54 3.77L7 21.96L1.5 20.5Z"/>`),
		g.Group(children),
	)
}