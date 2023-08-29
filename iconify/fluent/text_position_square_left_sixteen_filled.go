package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionSquareLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 1.75a.75.75 0 0 0 0 1.5h11a.75.75 0 0 0 0-1.5h-11Zm7.5 2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5H10ZM9.25 7.5a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5H10a.75.75 0 0 1-.75-.75ZM10 9.25a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5H10ZM1.75 12.5a.75.75 0 0 1 .75-.75h11a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1-.75-.75ZM3.25 7a1.75 1.75 0 1 1 3.5 0v3.5a.75.75 0 0 0 1.5 0V7a3.25 3.25 0 1 0-6.5 0v3.5a.75.75 0 0 0 1.5 0V7Z"/>`),
		g.Group(children),
	)
}