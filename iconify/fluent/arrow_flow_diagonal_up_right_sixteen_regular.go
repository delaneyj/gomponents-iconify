package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFlowDiagonalUpRightSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 3a.5.5 0 0 0 0 1h2.793L5.879 9.414a2.5 2.5 0 1 0 .707.707L12 4.707V7.5a.5.5 0 0 0 1 0v-4a.5.5 0 0 0-.5-.5h-4ZM3 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}