package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsCrossoverUpCrossMoveOverArrowArrowsRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5L.5.5m0 4v-4h4M9 5L13.5.5m0 4v-4h-4M5 9L.5 13.5"/>`),
		g.Group(children),
	)
}