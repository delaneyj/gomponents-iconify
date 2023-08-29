package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 128v1664h1664v128H128V128h128zm1536 128v256h-256V256h256zM768 384v256H512V384h256zm512 384v256h-256V768h256zm-512 512v256H512v-256h256zm896-128h256v256h-256v-256z"/>`),
		g.Group(children),
	)
}