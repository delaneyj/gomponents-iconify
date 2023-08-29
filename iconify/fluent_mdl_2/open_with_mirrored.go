package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenWithMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m0 1219l317-317l317 317l-90 90l-162-163v583H253v-583L90 1309l-90-90zm253-834h1153V256H253v129zm256 1280h897v-129H509v129zM1661 512h384V128h-384v384zm256-128h-128V256h128v128zm-256 768h384V768h-384v384zm256-128h-128V896h128v128zm-256 768h384v-384h-384v384zm256-128h-128v-128h128v128zM637 1024h768V896H509l128 128z"/>`),
		g.Group(children),
	)
}