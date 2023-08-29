package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableInsertColumnSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.5a.5.5 0 0 0-1 0v11a.5.5 0 0 0 1 0v-11Zm11 0a.5.5 0 0 0-1 0v11a.5.5 0 0 0 1 0v-11ZM6.5 2A1.5 1.5 0 0 0 5 3.5V5h6V3.5A1.5 1.5 0 0 0 9.5 2h-3ZM5 10V6h6v4H5Zm0 1h6v1.5A1.5 1.5 0 0 1 9.5 14h-3A1.5 1.5 0 0 1 5 12.5V11Z"/>`),
		g.Group(children),
	)
}