package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftDownTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.22 10.72a.75.75 0 0 0 0 1.06l4 4a.75.75 0 0 0 1.06 0l4-4a.75.75 0 1 0-1.06-1.06L8.5 13.44V7A1.5 1.5 0 0 1 10 5.5h6.25a.75.75 0 0 0 0-1.5H10a3 3 0 0 0-3 3v6.44l-2.72-2.72a.75.75 0 0 0-1.06 0Z"/>`),
		g.Group(children),
	)
}