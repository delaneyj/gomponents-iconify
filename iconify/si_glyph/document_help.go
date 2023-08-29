package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentHelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.969.062H3.031v15.92h10.938V4.062h-5v-4zm.047 12.972H7.954v-1.081h1.062v1.081zm1-7.05v1.004l1.001-.004v2.031h-1.001v.996h-1v1.011H7.969V8.968h2.016V7H7.016v1.031H5.969V6.969h1v-.984h3.047v-.001z"/><path d="M10.04.076v2.912H14L10.04.076z"/></g>`),
		g.Group(children),
	)
}