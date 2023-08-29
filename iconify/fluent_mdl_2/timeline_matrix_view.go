package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimelineMatrixView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 256v384h1920v128h-384v384h384v128h-384v384h384v128h-384v128h-128v-128H896v128H768v-128H128v128H0V128h2048v128H128zm640 1408v-384H128v384h640zm768-384H896v384h640v-384zm0-512H896v384h640V768zm-1408 0v384h640V768H128z"/>`),
		g.Group(children),
	)
}