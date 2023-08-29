package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnSingleCompareSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 1A1.75 1.75 0 0 0 3 2.75v10.5A1.75 1.75 0 0 0 4.75 15h6.5A1.75 1.75 0 0 0 13 13.25V2.75A1.75 1.75 0 0 0 11.25 1h-6.5ZM4 5h8v1H4V5Zm0 4h8v2H4V9Z"/>`),
		g.Group(children),
	)
}