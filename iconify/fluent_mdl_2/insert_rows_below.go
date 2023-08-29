package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertRowsBelow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1664h-640l128-128h384v-384h-768V768H768v512H128v384h256l128 128H0V128h2048zM640 256H128v384h512V256zm640 0H768v384h512V256zm640 0h-512v384h512V256zm-621 1139l90 90l-429 429l-429-429l90-90l275 275V896h128v774l275-275z"/>`),
		g.Group(children),
	)
}