package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 1664h256v128h-256v256h-128v-256h-256v-128h256v-256h128v256zM384 128v1536h768v128H256V0h859l549 549v731h-128V640h-512V128H384zm768 91v293h293l-293-293z"/>`),
		g.Group(children),
	)
}