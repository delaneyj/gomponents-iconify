package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageListMirroredSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 0v1792h1536V0H256zm1152 1152h128v128h-128v-128zm0-384h128v128h-128V768zm0-384h128v128h-128V384zm-896 768h768v128H512v-128zm0-384h768v128H512V768zm0-384h768v128H512V384z"/>`),
		g.Group(children),
	)
}