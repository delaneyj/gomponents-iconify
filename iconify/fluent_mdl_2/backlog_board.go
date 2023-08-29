package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacklogBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 384h512v384H256V384zm128 256h256V512H384v128zM256 896h512v384H256V896zm128 256h256v-128H384v128zm512-768h512v384H896V384zm128 256h256V512h-256v128zm512-256h512v384h-512V384zm128 256h256V512h-256v128zM896 896h512v384H896V896zm128 256h256v-128h-256v128zm512-256h512v384h-512V896zm128 256h256v-128h-256v128zM256 1408h512v384H256v-384zm128 256h256v-128H384v128zm1152-256h512v384h-512v-384zm128 256h256v-128h-256v128zm384-1536v128H128v1664H0V128h2048z"/>`),
		g.Group(children),
	)
}