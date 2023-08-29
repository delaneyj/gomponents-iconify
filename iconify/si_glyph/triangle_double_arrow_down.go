package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDoubleArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.044 16.565a1.49 1.49 0 0 1-2.104 0l-6.442-6.444c-.582-.581-.839-2.103 1-2.103h12.988c1.901 0 1.582 1.521 1.001 2.103l-6.443 6.444z"/><path d="M10.044 9.565a1.49 1.49 0 0 1-2.104 0L1.498 3.121c-.582-.581-.839-2.103 1-2.103h12.988c1.901 0 1.582 1.521 1.001 2.103l-6.443 6.444z"/></g>`),
		g.Group(children),
	)
}