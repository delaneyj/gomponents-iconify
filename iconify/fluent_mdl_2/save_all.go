package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 768q27 0 50 10t40 27t28 41t10 50v1152H933l-165-165V896q0-27 10-50t27-40t41-28t50-10h1024zm-768 128v384h512V896h-512zm512 1024v-256h-512v256h128v-128h128v128h256zm256-1024h-128v512h-768V896H896v933l91 91h37v-384h768v384h128V896zM256 640V128H128v933l91 91h37V768h384v128H384v256h128v-128h128v256H165L0 1115V128q0-27 10-50t27-40t41-28t50-10h1024q27 0 50 10t40 27t28 41t10 50v512h-128V128h-128v512H256zm128-128h512V128H384v384z"/>`),
		g.Group(children),
	)
}