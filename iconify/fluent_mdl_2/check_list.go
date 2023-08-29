package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 896V768h1408v128H640zm0-512h1408v128H640V384zm0 896v-128h1408v128H640zm0 384v-128h1408v128H640zM192 469l211-210l90 90l-301 301L19 477l90-90l83 82zm0 384l211-210l90 90l-301 301L19 861l90-90l83 82zm0 384l211-210l90 90l-301 301l-173-173l90-90l83 82zm0 384l211-210l90 90l-301 301l-173-173l90-90l83 82z"/>`),
		g.Group(children),
	)
}