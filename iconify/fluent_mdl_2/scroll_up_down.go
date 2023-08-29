package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 219L429 813l-90-90l685-686l685 686l-90 90l-595-594zm0 1610l595-594l90 90l-685 686l-685-686l90-90l595 594z"/>`),
		g.Group(children),
	)
}