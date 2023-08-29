package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 512h640v1152H0V512h640V256h768v256zM768 384v128h512V384H768zm1152 256H128v256h384V768h128v128h768V768h128v128h384V640zM128 1536h1792v-512h-384v128h-128v-128H640v128H512v-128H128v512z"/>`),
		g.Group(children),
	)
}