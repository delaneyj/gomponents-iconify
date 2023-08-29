package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M15 9.047v4H3v-4H1V16h15.969V9.047H15z"/><path d="M8.997 1L6 4.963l2.016.021v4.985h1.937V4.984h2L8.997 1z"/></g>`),
		g.Group(children),
	)
}