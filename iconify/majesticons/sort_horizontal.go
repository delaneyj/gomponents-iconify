package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M17.707 4.293A1 1 0 0 0 16 5v2h-6a1 1 0 1 0 0 2h6v2a1 1 0 0 0 1.707.707l3-3a1 1 0 0 0 0-1.414l-3-3zm-11.414 8A1 1 0 0 1 8 13v2h6a1 1 0 1 1 0 2H8v2a1 1 0 0 1-1.707.707l-3-3a1 1 0 0 1 0-1.414l3-3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}