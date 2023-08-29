package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFlowUpRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.854 2.146a.5.5 0 0 0-.708.708L13.293 5H9.5a2 2 0 0 0-2 2v2a1 1 0 0 1-1 1h-.55a2.5 2.5 0 1 0 0 1h.55a2 2 0 0 0 2-2V7a1 1 0 0 1 1-1h3.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3a.5.5 0 0 0 0-.708l-3-3Z"/>`),
		g.Group(children),
	)
}