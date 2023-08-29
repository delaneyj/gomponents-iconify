package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDoubleArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.446 6.773c.581.581.9 2.103-1.001 2.103H1.457c-1.839 0-1.582-1.521-1-2.103L6.898.329a1.49 1.49 0 0 1 2.104 0l6.444 6.444z"/><path d="M15.446 13.773c.581.581.9 2.103-1.001 2.103H1.457c-1.839 0-1.582-1.521-1-2.103l6.441-6.444a1.49 1.49 0 0 1 2.104 0l6.444 6.444z"/></g>`),
		g.Group(children),
	)
}