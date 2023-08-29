package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnowledgeArticle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H128V0h1115l549 549zm-512-37h293l-293-293v293zm384 128h-512V128H256v1792h1408V640zm-640 0H384V512h640v128zM384 768h1152v128H384V768zm0 256h1152v128H384v-128zm0 256h1152v128H384v-128zm0 256h1152v128H384v-128z"/>`),
		g.Group(children),
	)
}