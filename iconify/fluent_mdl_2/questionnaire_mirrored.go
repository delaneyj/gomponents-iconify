package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionnaireMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H128V0h1115l549 549zm-512-37h293l-293-293v293zm384 128h-512V128H256v1792h1408V640zm-512 256h384v384h-384V896zm128 256h128v-128h-128v128zm-896-128h640v128H384v-128zm0-512h640v128H384V512zm768 896h384v384h-384v-384zm128 256h128v-128h-128v128zm-896-128h640v128H384v-128z"/>`),
		g.Group(children),
	)
}