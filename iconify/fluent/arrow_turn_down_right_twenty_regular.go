package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnDownRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.146 16.854a.5.5 0 0 0 .708 0l4-4a.5.5 0 0 0 0-.708l-4-4a.5.5 0 0 0-.708.708L14.293 12H8a2 2 0 0 1-2-2V3.5a.5.5 0 0 0-1 0V10a3 3 0 0 0 3 3h6.293l-3.147 3.146a.5.5 0 0 0 0 .708Z"/>`),
		g.Group(children),
	)
}