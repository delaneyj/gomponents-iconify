package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRightTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.702 3.282C5.069 2.659 4 3.107 4 3.994v4.012c0 .887 1.07 1.335 1.702.713l2.037-2.006a1 1 0 0 0 0-1.425L5.702 3.282Z"/>`),
		g.Group(children),
	)
}