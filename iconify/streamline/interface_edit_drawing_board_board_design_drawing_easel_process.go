package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditDrawingBoardBoardDesignDrawingEaselProcess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 2.5h8a1 1 0 0 1 1 1V10h0H2h0V3.5a1 1 0 0 1 1-1ZM.5 10h13M7 2.5v-2M5.5 10L3 13.5M8.5 10l2.5 3.5"/>`),
		g.Group(children),
	)
}