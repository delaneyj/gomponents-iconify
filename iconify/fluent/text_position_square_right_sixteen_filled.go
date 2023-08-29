package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionSquareRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 1.75a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1 0-1.5h11ZM6 4.25a.75.75 0 0 1 0 1.5H2.5a.75.75 0 0 1 0-1.5H6Zm.75 3.25A.75.75 0 0 0 6 6.75H2.5a.75.75 0 0 0 0 1.5H6a.75.75 0 0 0 .75-.75ZM6 9.25a.75.75 0 0 1 0 1.5H2.5a.75.75 0 0 1 0-1.5H6Zm8.25 3.25a.75.75 0 0 0-.75-.75h-11a.75.75 0 0 0 0 1.5h11a.75.75 0 0 0 .75-.75ZM12.75 7a1.75 1.75 0 1 0-3.5 0v3.5a.75.75 0 0 1-1.5 0V7a3.25 3.25 0 1 1 6.5 0v3.5a.75.75 0 0 1-1.5 0V7Z"/>`),
		g.Group(children),
	)
}