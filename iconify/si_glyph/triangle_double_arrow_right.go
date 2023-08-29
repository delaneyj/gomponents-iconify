package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDoubleArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.113 15.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104l-6.444 6.443z"/><path d="M2.113 15.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104l-6.444 6.443z"/></g>`),
		g.Group(children),
	)
}