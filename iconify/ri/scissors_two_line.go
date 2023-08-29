package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13.41l-2.554 2.555a4 4 0 1 1-1.414-1.414l2.554-2.554l-6.021-6.021a2 2 0 0 1 0-2.829L12 10.582l7.435-7.435a2 2 0 0 1 0 2.829l-6.02 6.02l2.553 2.554a4 4 0 1 1-1.414 1.414L12 13.412Zm-6 6.587a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}