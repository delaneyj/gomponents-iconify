package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedListTwoMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384v128h1536V384H0zm1536 512H0v128h1536V896zm0 512H0v128h1536v-128zm512-1152h-384v384h384V256zm-128 256h-128V384h128v128zm128 256h-384v384h384V768zm-128 256h-128V896h128v128zm128 256h-384v384h384v-384zm-128 256h-128v-128h128v128z"/>`),
		g.Group(children),
	)
}