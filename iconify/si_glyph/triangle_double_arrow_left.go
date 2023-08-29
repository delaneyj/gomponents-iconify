package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDoubleArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.446 10.052a1.49 1.49 0 0 1 0-2.104L7.89 1.506c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001l-6.444-6.443z"/><path d="M8.446 10.052a1.49 1.49 0 0 1 0-2.104l6.444-6.442c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001l-6.444-6.443z"/></g>`),
		g.Group(children),
	)
}