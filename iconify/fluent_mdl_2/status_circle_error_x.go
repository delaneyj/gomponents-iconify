package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusCircleErrorX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1050 1088l371 371l-90 90l-371-371l-371 371l-90-90l371-371l-371-371l90-90l371 371l371-371l90 90l-371 371z"/>`),
		g.Group(children),
	)
}