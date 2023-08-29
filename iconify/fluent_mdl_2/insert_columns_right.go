package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertColumnsRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 640l-128-128V256h-512v512H896v384h512v512h512v-256l128-128v512H0V128h2048v512zM640 1280H128v384h512v-384zm0-512H128v384h512V768zm0-512H128v384h512V256zm883 1043l275-275h-774V896h774l-275-275l90-90l429 429l-429 429l-90-90z"/>`),
		g.Group(children),
	)
}