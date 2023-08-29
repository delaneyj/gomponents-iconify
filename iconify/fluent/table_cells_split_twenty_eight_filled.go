package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsSplitTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 7.5H25v-.75A3.75 3.75 0 0 0 21.25 3H14.5v4.5ZM13 3H6.75A3.75 3.75 0 0 0 3 6.75v.75h10V3Zm1.5 22h6.75A3.75 3.75 0 0 0 25 21.25v-.75H14.5V25ZM13 20.5H3v.75A3.75 3.75 0 0 0 6.75 25H13v-4.5ZM3 9h22v10H3V9Zm10 1v8h1.5v-8H13Z"/>`),
		g.Group(children),
	)
}