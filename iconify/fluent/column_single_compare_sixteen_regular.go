package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnSingleCompareSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0 1 11.25 15h-6.5A1.75 1.75 0 0 1 3 13.25V2.75ZM4.75 2a.75.75 0 0 0-.75.75V5h8V2.75a.75.75 0 0 0-.75-.75h-6.5ZM4 9h8V6H4v3Zm0 2v2.25c0 .414.336.75.75.75h6.5a.75.75 0 0 0 .75-.75V11H4Z"/>`),
		g.Group(children),
	)
}