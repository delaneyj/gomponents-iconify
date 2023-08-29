package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 384h1664v1024h-768v128h256v128H768v-128h256v-128H640v256H128V896h128V384zm256 1152v-512H256v512h256zm128-256h1152V512H384v384h256v384z"/>`),
		g.Group(children),
	)
}