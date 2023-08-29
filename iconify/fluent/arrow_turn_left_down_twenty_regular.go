package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.146 11.146a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708 0l4-4a.5.5 0 0 0-.708-.708L8 14.293V8a2 2 0 0 1 2-2h6.5a.5.5 0 0 0 0-1H10a3 3 0 0 0-3 3v6.293l-3.146-3.147a.5.5 0 0 0-.708 0Z"/>`),
		g.Group(children),
	)
}