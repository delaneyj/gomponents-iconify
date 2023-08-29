package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSymlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h640v128H128V0h1115l549 549v347h-128V640h-512V128H256v1792zM1280 512h293l-293-293v293zm768 512v1024H1024V1024h1024zm-128 128h-768v768h768v-768zm-621 531l274-275h-165v-128h384v384h-128v-165l-275 274l-90-90z"/>`),
		g.Group(children),
	)
}