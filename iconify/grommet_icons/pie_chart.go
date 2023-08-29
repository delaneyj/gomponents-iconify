package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10 23a9 9 0 0 1 0-18v9l1.162 1.162l5.202 5.202A8.972 8.972 0 0 1 10 23Zm4-13V1a9 9 0 0 1 9 9h-9Zm0 3h8a8.964 8.964 0 0 1-2.107 5.787L14 13Z"/>`),
		g.Group(children),
	)
}