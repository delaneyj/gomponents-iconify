package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExploreContentSingle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 256v1664H128V256h1664zm-128 128H256v1408h1408V384zm-768 768H512v-128h384V640h128v384h384v128h-384v384H896v-384z"/>`),
		g.Group(children),
	)
}