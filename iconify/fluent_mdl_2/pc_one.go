package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PCOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1536h-640v128h256v128h-640v-128h256v-128H896v256H0V128h896v512h1152v896zm-128-128V768H768v640h1152zM768 1664v-128H256v-128h384v-128H256v-128h384V640h128V256H128v1408h640zM640 384v128H256V384h384z"/>`),
		g.Group(children),
	)
}