package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsAlarmSnooze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M17.284 3.707l1.412-1.416l3.01 3l-1.413 1.417zm-10.586 0l-2.99 2.999L2.29 5.294l2.99-3zM12 4c-4.878 0-9 4.121-9 9s4.122 9 9 9c4.879 0 9-4.121 9-9s-4.121-9-9-9zm4 13H8.131l4-6H8V9h7.868l-1.035 1.554l-.001.001L11.869 15H16v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}