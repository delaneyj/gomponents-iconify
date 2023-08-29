package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionSquareRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.25 4a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1 0-1.5h16.5Zm-9.5 3.5a.75.75 0 0 1 0 1.5h-7a.75.75 0 1 1 0-1.5h7Zm.75 7.75a.75.75 0 0 0-.75-.75h-7a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 .75-.75ZM10.75 11a.75.75 0 0 1 0 1.5h-7a.75.75 0 0 1 0-1.5h7ZM21 18.75a.75.75 0 0 0-.75-.75H3.75a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 .75-.75ZM19.5 11a2.5 2.5 0 0 0-5 0v4.75a.75.75 0 0 1-1.5 0V11a4 4 0 0 1 8 0v4.75a.75.75 0 0 1-1.5 0V11Z"/>`),
		g.Group(children),
	)
}