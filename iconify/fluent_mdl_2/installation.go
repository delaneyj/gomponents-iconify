package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Installation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 1664h128v384H128v-384h128v256h1408v-256zm-147-531l-557 557l-557-557l90-90l403 402V128h128v1317l403-402l90 90z"/>`),
		g.Group(children),
	)
}