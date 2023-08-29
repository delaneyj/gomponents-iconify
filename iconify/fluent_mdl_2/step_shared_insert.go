package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepSharedInsert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 768v1152H896v-256H128v-640h512V640h512V256h768v512h128zM896 1536v-256h512V768h384V384h-512v384H768v384H256v384h640zm1024-640h-384v512h-512v384h896V896zM518 384H0V256h518L355 93l90-90l317 317l-317 317l-90-90l163-163z"/>`),
		g.Group(children),
	)
}