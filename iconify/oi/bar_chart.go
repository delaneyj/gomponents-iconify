package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v7h8V6H1V0H0zm5 0v5h2V0H5zM2 2v3h2V2H2z"/>`),
		g.Group(children),
	)
}