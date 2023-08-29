package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigateExternalInline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 1408h128v640H256V640h640v128H384v1152h1152v-512zm128-768v640h-128V859l-595 594l-90-90l594-595h-421V640h640z"/>`),
		g.Group(children),
	)
}