package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditBrushFourDesignRubberStampSuppliesTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 10.5h11v2a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-2h0Zm9-4a1 1 0 0 1-1-1V3a2.5 2.5 0 0 0-5 0v2.5a1 1 0 0 1-1 1h-1a2 2 0 0 0-2 2v1a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-1a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}