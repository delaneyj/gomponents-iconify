package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsVerticalScrollPointMoveScrollVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.13 3.48L7.26.61a.36.36 0 0 0-.52 0L3.87 3.48m6.26 7.04l-2.87 2.87a.36.36 0 0 1-.52 0l-2.87-2.87"/><circle cx="7" cy="7" r="1.25"/></g>`),
		g.Group(children),
	)
}