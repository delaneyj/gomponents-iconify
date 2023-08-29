package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.961 3.031h-.992v-.997h-1.016V.031h-.938v2.003h-.957v.997h-1.023v2H9V16h6.982V5.031h-1.021v-2zm-1.992.938h1v1h-1v-1zm-2 .015h1v1h-1v-1zM11 13.956h-1v-1h1v1zm0-1.999h-1v-1h1v1zm0-1.999h-1v-1h1v1zM11 8h-1V7h1v1zm2 5.956h-1v-1h1v1zm0-1.999h-1v-1h1v1zm0-1.999h-1v-1h1v1zM13 8h-1V7h1v1zm2 5.956h-1v-1h1v1zm0-1.999h-1v-1h1v1zm0-1.999h-1v-1h1v1zm0-2.009h-1v-1h1v1zm-8.039-4.9H5.924v-.997h-.949V.047h-.928v2.005H3.014v.997h-.979v1.998H1V16h6.982V5.047H6.961V3.049zm-2.004.912h1v1h-1v-1zM3 3.961h1v1H3v-1zm0 10.032H2v-1h1v1zm0-2.004H2v-1h1v1zm0-2.018H2v-1h1v1zM3 8H2V7h1v1zm2 5.969H4v-1h1v1zm0-1.995H4v-1h1v1zm0-1.993H4v-1h1v1zm0-2.006H4v-1h1v1zm2 5.984H6v-1h1v1zm0-1.998H6v-1h1v1zm0-1.998H6v-1h1v1zm0-2.008H6v-1h1v1z"/>`),
		g.Group(children),
	)
}