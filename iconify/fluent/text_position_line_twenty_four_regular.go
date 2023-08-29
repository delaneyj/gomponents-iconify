package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionLineTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 4a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5H3.75ZM7 8.5A2.5 2.5 0 0 0 4.5 11v4.75a.75.75 0 0 1-1.5 0V11a4 4 0 1 1 8 0v4.75a.75.75 0 0 1-1.5 0V11A2.5 2.5 0 0 0 7 8.5Zm6.75 6a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5h-6.5Zm-10 3.5a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5H3.75Z"/>`),
		g.Group(children),
	)
}