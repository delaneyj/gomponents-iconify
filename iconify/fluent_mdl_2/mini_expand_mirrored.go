package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniExpandMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m805 1024l91-91l-549-549h421V256H128v640h128V475l549 549zm-549 128H128v640h1792V256h-896v128h768v768h-768v512H256v-512zm1536 128v384h-640v-384h640z"/>`),
		g.Group(children),
	)
}