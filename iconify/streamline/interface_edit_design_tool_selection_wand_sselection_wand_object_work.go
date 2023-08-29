package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditDesignToolSelectionWandSselectionWandObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 13.5l9-9m2-2l1-1M9 2V.5M12 5h1.5M11 7l1 1M6 2l1 1"/>`),
		g.Group(children),
	)
}