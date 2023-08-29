package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedLineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m702 848l-405 355l660 330l963-602v138l-957 598l-707-353v479h1664v128H128V129h128v953l452-396l252 255l704-704l256 256v166l-256-256l-704 704l-258-259z"/>`),
		g.Group(children),
	)
}