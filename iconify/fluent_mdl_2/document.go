package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M549 0h1243v1755l-293 293H256V293L549 0zm1115 1701V128H603L384 347v1573h1061l219-219z"/>`),
		g.Group(children),
	)
}