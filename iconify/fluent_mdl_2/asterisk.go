package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1943 568l-791 456l791 456l-64 112l-791-457v913H960v-913l-791 457l-64-112l791-456l-791-456l64-112l791 457V0h128v913l791-457l64 112z"/>`),
		g.Group(children),
	)
}