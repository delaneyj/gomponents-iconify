package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.883 4.01L12 4.005a1 1 0 0 1 .993.883l.007.117v11.584l2.293-2.294a1 1 0 0 1 1.32-.084l.094.083a1 1 0 0 1 .084 1.32l-.084.095l-3.996 4a1 1 0 0 1-1.32.083l-.094-.083l-4.004-4a1 1 0 0 1 1.32-1.498l.094.083L11 16.583V5.004a1 1 0 0 1 .883-.992L12 4.004l-.117.007Z"/>`),
		g.Group(children),
	)
}