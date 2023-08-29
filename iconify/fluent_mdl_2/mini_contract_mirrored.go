package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniContractMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1280H128v512h1792V256h-768v128h640v768h-768v512H256v-384zm1536 0v384h-640v-384h640zM256 1024h640V384H768v421L219 256l-91 91l549 549H256v128z"/>`),
		g.Group(children),
	)
}