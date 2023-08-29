package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionLineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 3.75a1 1 0 0 0 0 2h16.5a1 1 0 1 0 0-2H3.75Zm3.25 5A2.25 2.25 0 0 0 4.75 11v4.75a1 1 0 1 1-2 0V11a4.25 4.25 0 0 1 8.5 0v4.75a1 1 0 1 1-2 0V11A2.25 2.25 0 0 0 7 8.75Zm6.75 5.5a1 1 0 1 0 0 2h6.5a1 1 0 1 0 0-2h-6.5Zm-10 3.5a1 1 0 1 0 0 2h16.5a1 1 0 1 0 0-2H3.75Z"/>`),
		g.Group(children),
	)
}