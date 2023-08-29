package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.031 9.016v4H3v-4H1V16h15.938V9.016h-1.907z"/><path d="m9.072 9.947l2.91-3.876l-2.014-.021V1.065H8.03V6.05h-2l3.042 3.897z"/></g>`),
		g.Group(children),
	)
}