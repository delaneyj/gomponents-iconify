package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedBarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1919 1792v128H128V128h128v256h1024v512H256v128h1536v512H256v256h1663zM768 512v256h384V512H768zm-512 0v256h384V512H256zm768 640v256h640v-256h-640zm-768 0v256h640v-256H256z"/>`),
		g.Group(children),
	)
}