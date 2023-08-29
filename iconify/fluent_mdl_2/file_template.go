package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1243 0l549 549v219h-128V640h-512V128H256v1792h768v128H128V0h1115zm37 512h293l-293-293v293zm640 1408h128v128h-128v-128zm-256 0h128v128h-128v-128zm-256 0h128v128h-128v-128zm256-1024h128v128h-128V896zm-256 0h128v128h-128V896zm-256 1024h128v128h-128v-128zm768-256h128v128h-128v-128zm-768 0h128v128h-128v-128zm768-256h128v128h-128v-128zm-768 0h128v128h-128v-128zm768-256h128v128h-128v-128zm-768 0h128v128h-128v-128zm896-256v128h-128V896h128zm-896 0h128v128h-128V896z"/>`),
		g.Group(children),
	)
}