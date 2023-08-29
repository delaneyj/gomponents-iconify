package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeAltF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14.858H1.674a1 1 0 0 1-.88-1.474l1.9-3.529a1 1 0 0 1-.785-1.512l4.234-7.056a1 1 0 0 1 1.714 0l4.234 7.056a1 1 0 0 1-.785 1.512l1.9 3.529a1 1 0 0 1-.88 1.474H8v6H6v-6z"/>`),
		g.Group(children),
	)
}