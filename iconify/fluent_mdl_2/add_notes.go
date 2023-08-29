package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v128H0V256zm0 640V768h2048v128H0zm0 512v-128h1280v128H0zm0 512v-128h1280v128H0zm1790-640v256h256v128h-256v256h-128v-256h-256v-128h256v-256h128z"/>`),
		g.Group(children),
	)
}