package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CRMReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 1536h-384V384h384v1152zM1280 512v896h128V512h-128zm-256 1024H640V640h384v896zM768 768v640h128V768H768zM128 0h1792v1920H128v-384H0v-128h128V640H0V512h128V0zm1664 1792V128H256v384h128v128H256v768h128v128H256v256h1536z"/>`),
		g.Group(children),
	)
}