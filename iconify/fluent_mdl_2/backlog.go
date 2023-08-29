package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backlog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 256h512v384H128V256zm128 256h256V384H256v128zM128 768h512v384H128V768zm128 256h256V896H256v128zm512-768h512v384H768V256zm128 256h256V384H896v128zm1024-256v384h-512V256h512zm-128 128h-256v128h256V384zM768 768h512v384H768V768zm128 256h256V896H896v128zm-768 256h512v384H128v-384zm128 256h256v-128H256v128z"/>`),
		g.Group(children),
	)
}