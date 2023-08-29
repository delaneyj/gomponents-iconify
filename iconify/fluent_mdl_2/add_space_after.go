package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSpaceAfter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 768V640h1792v128H128zm0-640h1792v128H128V128zm896 1414l163-163l90 90l-317 317l-317-317l90-90l163 163V896h128v646z"/>`),
		g.Group(children),
	)
}