package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 3H9V1H7v2H6l2 3l2-3zm4 10l-1.58-1.18l.78-1.82l-2-.23L11 7.8l-1.82.78L8 7L6.82 8.58L5 7.8l-.23 2l-1.97.2l.78 1.82L2 13H0v1h16v-1h-2zM4 13a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 12.993L4 13z"/>`),
		g.Group(children),
	)
}