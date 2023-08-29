package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceCursorArrowTwoMouseSelectCursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.15 5.45a.5.5 0 0 0 0-1L1.83.56A1 1 0 0 0 .56 1.83L4.5 13.16a.5.5 0 0 0 1 0l2-5.66Z"/>`),
		g.Group(children),
	)
}