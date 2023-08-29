package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortVerticalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.293 6.293a1 1 0 0 0 1.414 1.414L7 6.414V14a1 1 0 1 0 2 0V6.414l1.293 1.293a1 1 0 1 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0l-3 3zm8 11.414a1 1 0 0 1 1.414-1.414L15 17.586V10a1 1 0 1 1 2 0v7.586l1.293-1.293a1 1 0 0 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3z"/></g>`),
		g.Group(children),
	)
}