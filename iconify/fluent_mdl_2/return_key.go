package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReturnKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1152H475l466 467l-90 90l-621-621l621-621l90 90l-466 467h1189V384h128v768z"/>`),
		g.Group(children),
	)
}