package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsSplitTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 3H11v4H3v-.75A3.25 3.25 0 0 1 6.25 3ZM21 7v-.75A3.25 3.25 0 0 0 17.75 3H12.5v4H21Zm-8.5 14h5.25A3.25 3.25 0 0 0 21 17.75V17h-8.5v4ZM3 8.5v7h18v-7H3Zm9.5 1.5v4H11v-4h1.5ZM3 17.75V17h8v4H6.25A3.25 3.25 0 0 1 3 17.75Z"/>`),
		g.Group(children),
	)
}