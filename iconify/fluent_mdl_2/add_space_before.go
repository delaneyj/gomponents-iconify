package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSpaceBefore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1792v-128h1792v128H128zm0-640h1792v128H128v-128zm1149-451l-317 317l-317-317l90-90l163 163V128h128v646l163-163l90 90z"/>`),
		g.Group(children),
	)
}