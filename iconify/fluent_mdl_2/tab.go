package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 384v1664H512v-384H128V0h1408v384h384zM256 1536h1152V128H256v1408zM1792 512h-256v1152H640v256h1152V512z"/>`),
		g.Group(children),
	)
}