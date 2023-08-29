package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsDiagonalScrollPointOneMoveScroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="1.5"/><path d="M9 .5h4a.5.5 0 0 1 .5.5v4M5 13.5H1a.5.5 0 0 1-.5-.5V9"/></g>`),
		g.Group(children),
	)
}