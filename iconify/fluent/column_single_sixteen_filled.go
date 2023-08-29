package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnSingleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.75C4 2.784 4.784 2 5.75 2h3.5c.477 0 .91.19 1.225.5h.025v.025c.31.316.5.748.5 1.225v8.5A1.75 1.75 0 0 1 9.25 14h-3.5A1.75 1.75 0 0 1 4 12.25v-8.5Zm1 8.5c0 .414.336.75.75.75h3.5a.75.75 0 0 0 .75-.75V6H5v6.25Z"/>`),
		g.Group(children),
	)
}