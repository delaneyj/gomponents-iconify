package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v128H0V128zm0 1664v-128h2048v128H0zm0-768V896h2048v128H0zm512-384V512h1536v128H512zm0 768v-128h1536v128H512z"/>`),
		g.Group(children),
	)
}