package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditBrushThreeBrushColorColorsDesignPaintPaintingRollerRolling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="4" x="2.5" y=".5" rx="1"/><path d="M8 9.5v-1a2 2 0 0 0-2-2H2.5a2 2 0 0 1-2-2v-1a1 1 0 0 1 1-1h1m4 11v-3a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v3"/></g>`),
		g.Group(children),
	)
}