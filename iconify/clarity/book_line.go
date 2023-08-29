package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M10 5.2h18v1.55H10z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M29 8H9.86A1.89 1.89 0 0 1 8 6a2 2 0 0 1 1.86-2H29a1 1 0 0 0 0-2H9.86A4 4 0 0 0 6 6a4.14 4.14 0 0 0 0 .49a1 1 0 0 0 0 .24V30a4 4 0 0 0 3.86 4H29a1 1 0 0 0 1-1V9.07A1.07 1.07 0 0 0 29 8Zm-1 24H9.86A2 2 0 0 1 8 30V9.55a3.63 3.63 0 0 0 1.86.45H28Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}