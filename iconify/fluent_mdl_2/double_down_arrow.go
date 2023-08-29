package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleDownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m512 1798l261-261l90 90l-415 415l-415-415l90-90l261 261V0h128v1798zm1285-261l90 90l-415 415l-415-415l90-90l261 261V0h128v1798l261-261z"/>`),
		g.Group(children),
	)
}