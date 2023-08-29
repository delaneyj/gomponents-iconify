package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddEvent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 128h384v1792H0V128h384V0h128v128h1024V0h128v128zM384 256H128v256h1792V256h-256v128h-128V256H512v128H384V256zM128 1792h1792V640H128v1152zm960-1024v384h384v128h-384v384H960v-384H576v-128h384V768h128z"/>`),
		g.Group(children),
	)
}