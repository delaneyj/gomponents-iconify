package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRightUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.854 8.854a.5.5 0 0 0 0-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L12 5.707V12a2 2 0 0 1-2 2H3.5a.5.5 0 0 0 0 1H10a3 3 0 0 0 3-3V5.707l3.146 3.147a.5.5 0 0 0 .708 0Z"/>`),
		g.Group(children),
	)
}