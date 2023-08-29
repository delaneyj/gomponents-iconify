package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v128H0V128h2048zM0 1664h2048v128H0v-128zm0-768h2048v128H0V896zm0-384h2048v128H0V512zm0 768h2048v128H0v-128z"/>`),
		g.Group(children),
	)
}