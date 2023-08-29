package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModelingView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 896h-896V384H512v256h384v896H512v256h640v-640h896v896h-896v-128H384v-384H0V640h384V256h768V0h896v896zm-768-768v128h640V128h-640zm0 256v384h640V384h-640zm0 896v128h640v-128h-640zm0 256v384h640v-384h-640zm-512-128v-384H128v384h640zm0-640H128v128h640V768z"/>`),
		g.Group(children),
	)
}