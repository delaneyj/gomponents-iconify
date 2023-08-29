package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataFiveTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 12a.75.75 0 0 1 .743.64l.008.11v2.503a.75.75 0 0 1-1.493.11L4 15.254V12.75a.75.75 0 0 1 .75-.75Z"/>`),
		g.Group(children),
	)
}