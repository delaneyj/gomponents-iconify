package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnVerticalSection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 256h640v1536h-640V256zm512 1408V384h-384v1280h384zM256 384V256h256v128H256zm384 0V256h256v128H640zM256 1792v-128h256v128H256zm384 0v-128h256v128H640zM0 896V640h128v256H0zm0-384V256h128v256H0zm0 1152v-256h128v256H0zm0-384v-256h128v256H0zm1024-256V768h128v256h-128zm0-384V384h128v256h-128zm0 1152v-256h128v256h-128zm0-384v-256h128v256h-128z"/>`),
		g.Group(children),
	)
}