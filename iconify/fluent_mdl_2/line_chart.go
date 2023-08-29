package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m705 978l-449 449v366h1664v128H128V129h128v1132l449-449l255 257l704-704l256 256v166l-256-256l-704 704l-255-257z"/>`),
		g.Group(children),
	)
}