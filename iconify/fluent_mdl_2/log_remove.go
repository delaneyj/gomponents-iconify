package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1792h1092l-128 128H128v-384H0v-128h128V640H0V512h128V0h1792v1220l-128 128V128H256v384h128v128H256v768h128v128H256v256zm768-1152v896H640V640h384zM896 768H768v640h128V768zm512 510V512h-128v894l-95 95l34 35h-67V384h384v835l-35-34l-93 93zm411 450l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91l-226 226z"/>`),
		g.Group(children),
	)
}