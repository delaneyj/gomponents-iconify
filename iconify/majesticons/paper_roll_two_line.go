package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperRollTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 9a5 5 0 1 0 5-5m-5 5a5 5 0 0 1 5-5m-5 5v11l-1.5-1l-2 1l-2-1L4 20V9c0-4 3.333-5 5-5h7"/><circle cx="16" cy="9" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}