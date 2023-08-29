package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1243 1024l-91-91l549-549h-421V256h640v640h-128V475l-549 549zm549 128h128v640H128V256h896v128H256v768h768v512h768v-512zM256 1280v384h640v-384H256z"/>`),
		g.Group(children),
	)
}