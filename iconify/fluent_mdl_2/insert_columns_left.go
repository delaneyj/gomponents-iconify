package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertColumnsLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1664H0v-512l128 128v256h512v-512h512V768H640V256H128v256L0 640V128h2048zm-128 1152h-512v384h512v-384zm0-512h-512v384h512V768zm0-512h-512v384h512V256zM525 621L250 896h774v128H250l275 275l-90 90L6 960l429-429l90 90z"/>`),
		g.Group(children),
	)
}