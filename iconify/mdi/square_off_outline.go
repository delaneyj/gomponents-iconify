package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3L3 4.89V21h16.11l1.73 1.73l1.27-1.27M5 19V6.89L17.11 19H5M8.2 5l-2-2H21v14.8l-2-2V5H8.2Z"/>`),
		g.Group(children),
	)
}