package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeBracketsSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 17h-.667a2 2 0 0 1-2-2v-1.889C7.333 12.556 6 12 6 12s1.333-.556 1.333-1.111V9a2 2 0 0 1 2-2H10m4 10h.667a2 2 0 0 0 2-2v-1.889C16.667 12.556 18 12 18 12s-1.333-.556-1.333-1.111V9a2 2 0 0 0-2-2H14"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`),
		g.Group(children),
	)
}