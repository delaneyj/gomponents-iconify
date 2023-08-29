package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.146 8.854a.5.5 0 0 1 0-.708l4-4a.5.5 0 0 1 .708 0l4 4a.5.5 0 0 1-.708.708L8 5.707V12a2 2 0 0 0 2 2h6.5a.5.5 0 0 1 0 1H10a3 3 0 0 1-3-3V5.707L3.854 8.854a.5.5 0 0 1-.708 0Z"/>`),
		g.Group(children),
	)
}