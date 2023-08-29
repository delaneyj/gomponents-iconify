package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublishContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 347L493 749l-90-90l557-557l557 557l-90 90l-403-402v1317H896V347zm768 1317h128v384H128v-384h128v256h1408v-256z"/>`),
		g.Group(children),
	)
}