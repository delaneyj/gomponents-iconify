package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageListSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 0v1792H256V0h1536zM640 1152H512v128h128v-128zm0-384H512v128h128V768zm0-384H512v128h128V384zm896 768H768v128h768v-128zm0-384H768v128h768V768zm0-384H768v128h768V384z"/>`),
		g.Group(children),
	)
}