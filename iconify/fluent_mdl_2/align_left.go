package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v128H0V128h2048zm0 1664H0v-128h2048v128zm0-768H0V896h2048v128zm-512-384H0V512h1536v128zm0 768H0v-128h1536v128z"/>`),
		g.Group(children),
	)
}