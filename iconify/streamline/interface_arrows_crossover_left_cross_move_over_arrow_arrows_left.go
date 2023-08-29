package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsCrossoverLeftCrossMoveOverArrowArrowsLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5.5l-13 13m4 0h-4v-4M5 5L.5.5m4 0h-4v4M9 9l4.5 4.5"/>`),
		g.Group(children),
	)
}