package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsSplitSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4h5.95a2.5 2.5 0 0 0-2.45-2H8v2ZM7 2H4.5a2.5 2.5 0 0 0-2.45 2H7V2Zm1 12h3.5a2.5 2.5 0 0 0 2.45-2H8v2Zm-1-2H2.05a2.5 2.5 0 0 0 2.45 2H7v-2Zm-5-1V5h12v6H2Zm5-5v4h1V6H7Z"/>`),
		g.Group(children),
	)
}