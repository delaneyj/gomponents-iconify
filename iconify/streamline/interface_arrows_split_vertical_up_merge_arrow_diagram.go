package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsSplitVerticalUpMergeArrowDiagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 3.5l-3-3l-3 3m7 10a4 4 0 0 1-4-4a4 4 0 0 1-4 4m4-4v-9"/>`),
		g.Group(children),
	)
}