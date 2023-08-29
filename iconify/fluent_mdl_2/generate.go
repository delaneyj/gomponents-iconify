package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Generate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2011 1728l-318 317l-90-90l163-163h-614v-128h614l-163-163l90-90l318 317zm-669 192l128 128H128V0h1115l549 549v734l-128-128V640h-512V128H256v1792h1086zm-62-1408h293l-293-293v293z"/>`),
		g.Group(children),
	)
}