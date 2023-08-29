package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 1664v-384h128v512h-512v-128h384zM1280 256h512v512h-128V384h-384V256zM256 768V256h512v128H384v384H256zm128 512v384h384v128H256v-512h128z"/>`),
		g.Group(children),
	)
}