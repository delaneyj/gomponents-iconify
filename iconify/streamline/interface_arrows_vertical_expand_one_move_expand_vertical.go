package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsVerticalExpandOneMoveExpandVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.25 5.1L6.74.61a.36.36 0 0 1 .52 0l4.49 4.49m-9.5 3.8l4.49 4.49a.36.36 0 0 0 .52 0l4.49-4.49"/>`),
		g.Group(children),
	)
}