package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCrosshair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1088 870l192 191v182l-192-193l-166 166l65 64H805L320 794L128 987v549h640v128H0V256h1792v1024h-128V384H128v421l192-191l512 512l256-256zm384-358q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19zm-64 384h128v384h-128V896zm0 768h128v384h-128v-384zm-512-256h384v128H896v-128zm1152 0v128h-384v-128h384z"/>`),
		g.Group(children),
	)
}