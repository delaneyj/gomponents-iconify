package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniContract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1280h128v512H128V256h768v128H256v768h768v512h768v-384zm-1536 0v384h640v-384H256zm1536-256h-640V384h128v421l549-549l91 91l-549 549h421v128z"/>`),
		g.Group(children),
	)
}