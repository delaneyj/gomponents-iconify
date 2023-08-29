package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v1664h-384v384H0V384h384V0h1664zm-128 1536V128H512v256h256v128H128v1408h1408v-640h128v256h256zm-979-339l-90-90l594-595h-421V384h640v640h-128V603l-595 594z"/>`),
		g.Group(children),
	)
}