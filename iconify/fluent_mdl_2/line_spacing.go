package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m448 102l365 365l-90 90l-211-210v549H384V347L173 557l-90-90l365-365zm275 1389l90 90l-365 365l-365-365l90-90l211 210v-549h128v549l211-210zM2048 384v128H1024V384h1024zM1024 768h1024v128H1024V768zm0 384h1024v128H1024v-128zm0 384h1024v128H1024v-128z"/>`),
		g.Group(children),
	)
}