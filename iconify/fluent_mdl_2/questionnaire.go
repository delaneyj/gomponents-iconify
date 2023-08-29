package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Questionnaire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H128V0h1115l549 549zm-512-37h293l-293-293v293zm384 128h-512V128H256v1792h1408V640zM384 896h384v384H384V896zm128 256h128v-128H512v128zM384 384h384v384H384V384zm128 256h128V512H512v128zm384 384h640v128H896v-128zm-512 384h384v384H384v-384zm128 256h128v-128H512v128zm384-128h640v128H896v-128z"/>`),
		g.Group(children),
	)
}