package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHorizontalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17.707 4.293a1 1 0 1 0-1.414 1.414L17.586 7H10a1 1 0 0 0 0 2h7.586l-1.293 1.293a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414l-3-3zm-11.414 8a1 1 0 0 1 1.414 1.414L6.414 15H14a1 1 0 1 1 0 2H6.414l1.293 1.293a1 1 0 1 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3z"/></g>`),
		g.Group(children),
	)
}