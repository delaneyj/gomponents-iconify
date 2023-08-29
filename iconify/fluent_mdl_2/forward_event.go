package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardEvent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 128h384v1297l-128-128V640H128v1152h1291l128 128H0V128h384V0h128v128h1024V0h128v128zM128 512h1792V256h-256v128h-128V256H512v128H384V256H128v256zm1507 861l90-90l317 317l-317 317l-90-90l163-163h-518v-128h518l-163-163z"/>`),
		g.Group(children),
	)
}