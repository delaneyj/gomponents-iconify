package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsShrinkVerticalMoveVerticalShrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.25 13.5L6.74 9a.37.37 0 0 1 .52 0l4.49 4.49M2.25.5L6.74 5a.37.37 0 0 0 .52 0L11.75.5"/>`),
		g.Group(children),
	)
}