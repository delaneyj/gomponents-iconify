package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1115 1024l914 915l-90 90l-915-914l-915 914l-90-90l914-915L19 109l90-90l915 914l915-914l90 90l-914 915z"/>`),
		g.Group(children),
	)
}