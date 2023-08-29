package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionSquareRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.25 3.75a1 1 0 1 1 0 2H3.75a1 1 0 0 1 0-2h16.5Zm-9.5 3.5a1 1 0 1 1 0 2h-7a1 1 0 0 1 0-2h7Zm1 8a1 1 0 0 0-1-1h-7a1 1 0 1 0 0 2h7a1 1 0 0 0 1-1Zm-1-4.5a1 1 0 1 1 0 2h-7a1 1 0 1 1 0-2h7Zm10.5 8a1 1 0 0 0-1-1H3.75a1 1 0 1 0 0 2h16.5a1 1 0 0 0 1-1Zm-2-7.75a2.25 2.25 0 0 0-4.5 0v4.75a1 1 0 1 1-2 0V11a4.25 4.25 0 0 1 8.5 0v4.75a1 1 0 1 1-2 0V11Z"/>`),
		g.Group(children),
	)
}