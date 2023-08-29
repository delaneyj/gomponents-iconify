package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCSS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 549v1499H128V0h1115l549 549zm-512-37h293l-293-293v293zm384 1408V640h-512V128H256v1792h1408zM1472 896h64v896h-128v-256H987l-256 256H576l896-896zm-64 219l-293 293h293v-293zM384 256h128v128H384V256zm0 256h128v128H384V512zm0 256h128v128H384V768z"/>`),
		g.Group(children),
	)
}