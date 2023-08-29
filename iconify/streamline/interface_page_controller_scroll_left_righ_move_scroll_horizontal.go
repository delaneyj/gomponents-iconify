package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePageControllerScrollLeftRighMoveScrollHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9 2.25l4.39 4.49a.36.36 0 0 1 0 .52L9 11.75m-4-9.5L.61 6.74a.36.36 0 0 0 0 .52L5 11.75"/>`),
		g.Group(children),
	)
}