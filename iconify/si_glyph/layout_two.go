package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 16H1.021V.042H17V16zm-1.042-.957V1h-14v14.043h14z"/><path d="M3 2h5v12H3zm7 0h5v12h-5z"/></g>`),
		g.Group(children),
	)
}