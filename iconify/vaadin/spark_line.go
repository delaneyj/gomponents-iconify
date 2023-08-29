package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparkLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 6a2 2 0 0 0-2 2v.16l-.81.25l-2.3-3.48l-1.73 4.32L6 3.44L3.7 8.22L2.06 6.91L0 8v1.08l1.94-1l2.11 1.7l1.56-3.22l1.23 6.19l2.27-5.68l1.68 2.52l1.55-.48A2 2 0 1 0 14.004 6H14zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}