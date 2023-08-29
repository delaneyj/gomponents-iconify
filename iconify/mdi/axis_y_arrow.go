package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisYArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.5 20.5L17 21.96l1.53-2.61L12 15.58L2.61 21l-1-1.73L11 13.85V3h2v10.85l6.53 3.76L21 15.03l1.5 5.47Z"/>`),
		g.Group(children),
	)
}