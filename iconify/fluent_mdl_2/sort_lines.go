package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortLines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m384 1638l163-163l90 90l-317 318L3 1565l90-90l163 163V128h128v1510zm384-358v-128h512v128H768zm0-384V768h896v128H768zm0-512h1280v128H768V384z"/>`),
		g.Group(children),
	)
}