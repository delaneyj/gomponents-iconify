package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowResults(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 1792v-128h1664v128H384zM0 384V256h128v128H0zm0 1024v-128h128v128H0zm384 0v-128h1664v128H384zm0-1152h1664v128H384V256zm0 512V640h1664v128H384z"/>`),
		g.Group(children),
	)
}